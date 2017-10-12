package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	// Caltech Library Package
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/pkgassets"
)

var (
	// Standard Options
	showHelp     bool
	showLicense  bool
	showVersion  bool
	showExamples bool
	outFName     string

	// App Options
	packageName  string
	commentFName string
	stripPrefix  string
	stripSuffix  string
)

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showExamples, "example", false, "display example(s)")
	flag.StringVar(&outFName, "o", "", "output filename")
	flag.StringVar(&outFName, "output", "", "output filename")

	// App Options
	flag.StringVar(&packageName, "p", "", "package name, if missing defauls to lowercase of variable name")
	flag.StringVar(&packageName, "package", "", "package name, if missing defauls to lowercase of variable name")
	flag.StringVar(&commentFName, "c", "", "comment file to be included")
	flag.StringVar(&commentFName, "comment", "", "comment file to be included")
	flag.StringVar(&stripPrefix, "strip-prefix", "", "strip the prefix from the map key")
	flag.StringVar(&stripSuffix, "strip-suffix", "", "strip the suffix from the map key")
}

func main() {
	var (
		err error
	)
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	cfg := cli.New(appName, "PKGASSETS", pkgassets.Version)
	cfg.LicenseText = fmt.Sprintf(pkgassets.LicenseText, appName, pkgassets.Version)
	cfg.UsageText = fmt.Sprintf("%s", Help["usage"])
	cfg.DescriptionText = fmt.Sprintf("%s", Help["description"])
	cfg.OptionText = "## OPTIONS \n\n"
	cfg.ExampleText = fmt.Sprintf("%s", Help["example"])

	// map in our help and examples
	for k, v := range Help {
		cfg.AddHelp(k, fmt.Sprintf("%s", v))
	}
	for k, v := range Examples {
		cfg.AddExample(k, fmt.Sprintf("%s", v))
	}

	// Process flags and update the environment as needed.
	if showHelp == true {
		if len(args) > 0 {
			fmt.Println(cfg.Help(args...))
		} else {
			fmt.Println(cfg.Usage())
		}
		os.Exit(0)
	}

	if showExamples == true {
		/*
			if len(args) > 0 {
				fmt.Println(cfg.Example(args...))
			} else {
				fmt.Println(cfg.ExampleText)
			}
		*/
		fmt.Println(cfg.Example(args...))
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, cfg.Usage())
		fmt.Fprintf(os.Stderr, "\n")
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "Missing map variable name\n")
		}
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Missing asset directory name\n")
		}
		os.Exit(1)
	}
	if (len(args) % 2) != 0 {
		fmt.Fprintf(os.Stderr, "Missing variable/assets directory names must be paired\n")
		fmt.Fprintf(os.Stderr, "Paramters provided, %q\n", strings.Join(args, ", "))
		os.Exit(1)
	}

	// Write out the Go source file.
	fp, err := cli.Create(outFName, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't write %s, %s\n", outFName, err)
		os.Exit(1)
	}
	defer fp.Close()

	// For each pair of mapVName/assetDir add a map to outFName
	for i := 0; (i + 1) < len(args); i += 2 {
		mapVName, assetDir := args[i], args[i+1]

		if packageName == "" {
			packageName = strings.ToLower(mapVName)
		}
		if outFName == "" {
			outFName = strings.ToLower(packageName) + ".go"
		}
		commentSrc := []byte{}
		if commentFName != "" {
			commentSrc, err = ioutil.ReadFile(commentFName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't read %s, %s\n", commentFName, err)
				os.Exit(1)
			}
		}

		if len(commentSrc) > 0 {
			fmt.Fprintf(fp, "\n//\n// %s\n//\n", bytes.Replace(commentSrc, []byte("\n"), []byte("\n// "), -1))
		}
		if i == 0 {
			fmt.Fprintf(fp, "package %s\n\nvar (\n\n", packageName)
		}
		fmt.Fprintf(fp, `    // %s is a map to asset files associated with %s package
    %s = map[string][]byte{`, mapVName, packageName, mapVName)
		// Walk the asset directory structure and for each file found at it to the map...
		if err = filepath.Walk(assetDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			fpath := strings.TrimPrefix(path, assetDir)
			if stripPrefix != "" {
				fpath = strings.TrimPrefix(fpath, stripPrefix)
			}
			if stripSuffix != "" {
				fpath = strings.TrimSuffix(fpath, stripSuffix)
			}
			bArray, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't read %q, %s", path, err)
				return nil
			}
			if bSrc, err := pkgassets.ByteArrayToDecl(bArray); err == nil {
				fmt.Fprintf(fp, "\n    %q: %s,\n", fpath, bSrc)
			} else {
				fmt.Fprintf(os.Stderr, "Can't convert to byte array notation %s, %s\n", path, err)
			}
			return nil
		}); err != nil {
			fmt.Fprintf(os.Stderr, "Can't walk path %q, %s\n", assetDir, err)
		}
		fmt.Fprintf(fp, `
	}
`)
	}
	fmt.Fprintf(fp, `
)

`)

}

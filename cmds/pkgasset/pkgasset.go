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
	"github.com/caltechlibrary/pkgasset"
)

var (
	usage = `USAGE: %s PACKAGE_NAME DIR_HOLDING_ASSETS`

	description = `
SYNOPSIS

%s generates a Go source file hosting a map of path and byte array
of file content harvested from directory holding the assets. The
path key starts with a slash and does not include the hosting directory
name (e.g. htdocs/index.html would become /index.html if htdocs was
used to harvest assets).
`

	examples = `
EXAMPLE

> %s MAP_VARAIBLE_NAME NAME_OF_DIRECTORY_HOLDING_ASSETS

This will result in a Go of type map[string][]byte holding the assets discovered by walking the directory
tree provided. The map's key will represent a path (beginning with "/") pointing at the asset ingested.

> %s DefaultSite htdocs

Assuming that _htdocs_ held

+ index.html
+ css/site.css

In this example the htdocs directory will be crawled and all the files found harvested as a an asset. The
path in the map will not include htdocs and would result in a Go source file like

    package defaultsite

    var DefaultSite = map[string][]byte{
        "/index.html": []byte("... the contents of index.html would be here ..."),
        "/css/site.css": []byte("... the contents of css/site.css would be here ..."),
    }

If a package name is not provided then the package name will a lowercase name of the map variable name (e.g. 
"var DefaultSite" becomes "package defaultsite"). Likewise if a output name is not provided then the file
name will be the name of the package plus the ".go" extension.
`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	outFName    string

	// App Options
	packageName  string
	commentFName string
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.StringVar(&outFName, "o", "", "output filename")
	flag.StringVar(&outFName, "output", "", "output filename")
	flag.StringVar(&packageName, "p", "", "package name, if missing defauls to lowercase of variable name")
	flag.StringVar(&packageName, "package", "", "package name, if missing defauls to lowercase of variable name")
	flag.StringVar(&commentFName, "c", "", "comment file to be included")
	flag.StringVar(&commentFName, "comment", "", "comment file to be included")
}

func main() {
	var (
		err error
	)
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	cfg := cli.New(appName, "PKGASSET", fmt.Sprintf(pkgasset.LicenseText, appName, pkgasset.Version), pkgasset.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName)

	// Process flags and update the environment as needed.
	if showHelp == true {
		fmt.Println(cfg.Usage())
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

	if len(args) != 2 {
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
	mapVName, assetDir := args[0], args[1]

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

	// Write out the Go source file.
	fp, err := os.Create(outFName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't write %s, %s\n", outFName, err)
		os.Exit(1)
	}
	defer fp.Close()

	if len(commentSrc) > 0 {
		fmt.Fprintf(fp, "//\n// %s\n//\n", bytes.Replace(commentSrc, []byte("\n"), []byte("\n// "), -1))
	}
	fmt.Fprintf(fp, "package %s\n\n", packageName)
	fmt.Fprintf(fp, `var (
    %s = map[string][]byte{`, mapVName)
	// Walk the asset directory structure and for each file found at it to the map...
	if err = filepath.Walk(assetDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fpath := strings.TrimPrefix(path, assetDir)
		bArray, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't read %q, %s", path, err)
			return nil
		}
		fmt.Fprintf(fp, "\n    %q: []byte(`%x`),", fpath, bArray)
		return nil
	}); err != nil {
		fmt.Fprintf(os.Stderr, "Can't walk path %q, %s\n", assetDir, err)
	}
	fmt.Fprintf(fp, `
	}
)`)

}

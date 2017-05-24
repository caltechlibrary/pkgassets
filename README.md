
# pkgassets

Generate a Go source file with embedded file assets. This is useful where you want to embed web content,
template source code and other assets that can be used for default behavior in a Go command line program
or service.

## USAGE

```
    pkgassets MAP_VAR_NAME NAME_OF_DIRECTORY_HOLDING_ASSETS
```

This will result in a Go map[string][]byte structure holding the assets discovered by walking the directory
tree provided. The map's key will represent a path (beginning with "/") pointing at the asset ingested.

```shell
    pkgassets DefaultSite htdocs
```

Assuming that _htdocs_ held

+ index.html
+ css/site.css

In this example the htdocs directory will be crawled and all the files found harvested as a an asset. The
path in the map will not include htdocs and would result in a Go source file, _defaultsite.go_ would like

```
    package defaultsite

    var DefaultSite = map[string][]byte{
        "/index.html": []byte("... the contents of index.html would be here ..."),
        "/css/site.css": []byte("... the contents of css/site.css would be here ..."),
    }
```

If a package name is not provided then the package name will be assumed to be the name of the file. The
Asset option will also be assume to be the name of the file.

## OPTIONS

+ -h, -help will display the help page
+ -l, -license will display license info
+ -v, -version will display version info
+ -o, -output will set the name of the file output otherwise it defaults to package name in lowercase plus the ".go" extension
+ -p, -package will set the name of the package if missing it defaults to a lower case version of map variable name
+ -c, -comment will include a file as a comment in the head of the source file created (e.g. for your copyright statement)





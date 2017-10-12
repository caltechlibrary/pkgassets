
## Description

pkgassets generates a Go source file hosting a map of path and byte array
of file content harvested from directory holding the assets. The
path key starts with a slash and does not include the hosting directory
name (e.g. htdocs/index.html would become /index.html if htdocs was
used to harvest assets).

#
# Simple Makefile
#

PROJECT = pkgassets

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\` -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

build: 
	echo "Bookstrapping bin/pkgassets"
	if [ ! -f cmd/pkgassets/help.go ]; then echo 'package main;var Help map[string][]byte' > cmd/pkgassets/help.go;fi
	if [ ! -f cmd/pkgassets/examples.go ]; then echo 'package main;var Examples map[string][]byte' > cmd/pkgassets/examples.go;fi
	go build -o bin/pkgassets cmd/pkgassets/pkgassets.go cmd/pkgassets/help.go cmd/pkgassets/examples.go
	echo "Bootstrap completed"
	if [ -f bin/pkgassets ]; then bin/pkgassets -o cmd/pkgassets/help.go -p main -ext=".md" -strip-prefix="/" -strip-suffix=".md" Help docs; fi
	if [ -f bin/pkgassets.exe ]; then bin/pkgassets.exe -o cmd/pkgassets/help.go -p main -ext=".md" -strip-prefix="/" -strip-suffix=".md" Help docs; fi
	if [ -f bin/pkgassets ]; then bin/pkgassets -o cmd/pkgassets/examples.go -p main -ext=".md" -strip-prefix="/" -strip-suffix=".md" Examples examples; fi
	if [ -f bin/pkgassets.exe ]; then bin/pkgassets.exe -o cmd/pkgassets/examples.go -p main -ext=".md" -strip-prefix="/" -strip-suffix=".md" Examples examples; fi
	git add cmd/pkgassets/help.go cmd/pkgassets/examples.go
	go build -o bin/pkgassets cmd/pkgassets/pkgassets.go cmd/pkgassets/help.go cmd/pkgassets/examples.go

lint:
	golint pkgassets.go
	golint cmd/pkgassets/pkgassets.go

format:
	gofmt -w pkgassets.go
	gofmt -w cmd/pkgassets/pkgassets.go

test:
	go test

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

clean:
	if [ -f bin/pkgassets ] && [ -f cmd/pkgassets/help.go ]; then rm cmd/pkgassets/help.go; fi
	if [ -f bin/pkgassets.exe ] && [ -f cmd/pkgassets/help.go ]; then rm cmd/pkgassets/help.go; fi
	if [ -f bin/pkgassets ] && [ -f cmd/pkgassets/examples.go ]; then rm cmd/pkgassets/examples.go; fi
	if [ -f bin/pkgassets.exe ] && [ -f cmd/pkgassets/examples.go ]; then rm cmd/pkgassets/examples.go; fi
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

man: build
	mkdir -p man/man1
	bin/pkgassets -generate-manpage | nroff -Tutf8 -man > man/man1/pkgassets.1


install: build
	env GOBIN=$(GOPATH)/bin go install cmd/pkgassets/pkgassets.go cmd/pkgassets/help.go cmd/pkgassets/examples.go


dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/pkgassets cmd/pkgassets/pkgassets.go cmd/pkgassets/help.go cmd/pkgassets/examples.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/pkgassets.exe cmd/pkgassets/pkgassets.go cmd/pkgassets/help.go cmd/pkgassets/examples.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/pkgassets cmd/pkgassets/pkgassets.go cmd/pkgassets/help.go cmd/pkgassets/examples.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/pkgassets cmd/pkgassets/pkgassets.go cmd/pkgassets/help.go cmd/pkgassets/examples.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-amd7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	./package-versions.bash > dist/package-versions.txt

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash


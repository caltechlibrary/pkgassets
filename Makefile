#
# Simple Makefile
#

PROJECT = pkgassets

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

build: 
	echo "Bookstrapping bin/pkgassets"
	if [ ! -f cmds/pkgassets/help.go ]; then echo 'package main;var Help map[string][]byte' > cmds/pkgassets/help.go;fi
	if [ ! -f cmds/pkgassets/examples.go ]; then echo 'package main;var Examples map[string][]byte' > cmds/pkgassets/examples.go;fi
	go build -o bin/pkgassets cmds/pkgassets/*.go
	echo "Bootstrap completed"
	if [ -f bin/pkgassets ]; then bin/pkgassets -o cmds/pkgassets/help.go -p main -strip-prefix="/" -strip-suffix=".md" Help docs; fi
	if [ -f bin/pkgassets.exe ]; then bin/pkgassets.exe -o cmds/pkgassets/help.go -p main -strip-prefix="/" -strip-suffix=".md" Help docs; fi
	if [ -f bin/pkgassets ]; then bin/pkgassets -o cmds/pkgassets/examples.go -p main -strip-prefix="/" -strip-suffix=".md" Examples examples; fi
	if [ -f bin/pkgassets.exe ]; then bin/pkgassets.exe -o cmds/pkgassets/examples.go -p main -strip-prefix="/" -strip-suffix=".md" Examples examples; fi
	go build -o bin/pkgassets cmds/pkgassets/pkgassets.go cmds/pkgassets/help.go cmds/pkgassets/examples.go

lint:
	golint pkgassets.go
	golint cmds/pkgassets/pkgassets.go

format:
	gofmt -w pkgassets.go
	gofmt -w cmds/pkgassets/pkgassets.go

test:
	go test

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

clean:
	if [ -f bin/pkgassets ] && [ -f cmds/pkgassets/help.go ]; then rm cmds/pkgassets/help.go; fi
	if [ -f bin/pkgassets.exe ] && [ -f cmds/pkgassets/help.go ]; then rm cmds/pkgassets/help.go; fi
	if [ -f bin/pkgassets ] && [ -f cmds/pkgassets/examples.go ]; then rm cmds/pkgassets/examples.go; fi
	if [ -f bin/pkgassets.exe ] && [ -f cmds/pkgassets/examples.go ]; then rm cmds/pkgassets/examples.go; fi
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

install:
	env GOBIN=$(GOPATH)/bin go install cmds/pkgassets/pkgassets.go cmds/pkgassets/help.go cmds/pkgassets/examples.go


dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/pkgassets cmds/pkgassets/*.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/pkgassets.exe cmds/pkgassets/*.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/pkgassets cmds/pkgassets/*.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/pkgassets cmds/pkgassets/*.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-amd7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash


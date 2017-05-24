#
# Simple Makefile
#

PROJECT = pkgassets

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

build: bin/pkgassets

bin/pkgassets: pkgassets.go cmds/pkgassets/pkgassets.go
	go build -o bin/pkgassets cmds/pkgassets/pkgassets.go

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
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then rm -f $(PROJECT)-$(VERSION)-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/pkgassets/pkgassets.go


dist/linux-amd64:
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/pkgassets cmds/pkgassets/pkgassets.go

dist/windows-amd64:
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/pkgassets.exe cmds/pkgassets/pkgassets.go

dist/macosx-amd64:
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/pkgassets cmds/pkgassets/pkgassets.go

dist/raspbian-arm7:
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/pkgassets cmds/pkgassets/pkgassets.go

release: dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	zip -r $(PROJECT)-$(VERSION)-release.zip dist/*

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash


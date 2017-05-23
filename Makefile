#
# Simple Makefile
#

PROJECT = pkgasset

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\" -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

build: bin/pkgasset

bin/pkgasset: pkgasset.go cmds/pkgasset/pkgasset.go
	go build -o bin/pkgasset cmds/pkgasset/pkgasset.go

lint:
	golint pkgasset.go
	golint cmds/pkgasset/pkgasset.go

format:
	gofmt -w pkgasset.go
	gofmt -w cmds/pkgasset/pkgasset.go

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
	env GOBIN=$(HOME)/bin go install cmds/pkgasset/pkgasset.go


dist/linux-amd64:
	env  GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/pkgasset cmds/pkgasset/pkgasset.go

dist/windows-amd64:
	env  GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/pkgasset.exe cmds/pkgasset/pkgasset.go

dist/macosx-amd64:
	env  GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/pkgasset cmds/pkgasset/pkgasset.go

dist/raspbian-arm7:
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/pkgasset cmds/pkgasset/pkgasset.go

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


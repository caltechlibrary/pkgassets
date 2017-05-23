#!/bin/bash

TITLE="pkgasset: an asset packager and project setup tool"

START=$(pwd)
cd "$(dirname "$0")"

function checkApp() {
	APP_NAME="$(which "$1")"
	if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$1" ]; then
		echo "Missing $APP_NAME"
		exit 1
	fi
}

function softwareCheck() {
	for APP_NAME in "$@"; do
		checkApp "$APP_NAME"
	done
}

function MakePage() {
	nav="$1"
	content="$2"
	html="$3"
	echo "Rendering $html"
	mkpage \
		"title=text:${TITLE}" \
		"nav=$nav" \
		"content=$content" \
		"sitebuilt=text:Updated $(date)" \
		"copyright=copyright.md" \
		page.tmpl >"$html"
}

echo "Checking necessary software is installed"
softwareCheck mkpage
echo "Generating website index.html"
MakePage nav.md README.md index.html
echo "Generating install.html"
MakePage nav.md INSTALL.md install.html
echo "Generating license.html"
MakePage nav.md "markdown:$(cat LICENSE)" license.html

cd "$START"

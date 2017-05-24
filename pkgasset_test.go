package pkgassets

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTextFile(t *testing.T) {
	txt, err := ioutil.ReadFile("README.md")
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	if src, err := ByteArrayToDecl(txt); err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	} else {
		fmt.Printf("src = %s\n", src)
	}
}

func TestBinFile(t *testing.T) {
	t.Errorf("TestBinFile() not implemented")
}

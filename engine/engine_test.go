package engine

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_parseInValue(t *testing.T) {
	content := []byte("abc\n")
	tmpFile, err := ioutil.TempFile("", "stdin")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpFile.Name()) // clean up

	if _, err := tmpFile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpFile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	os.Stdin = tmpFile
	result, err2 := parseInValue()
	if err2 != nil || result != "abc\n" {
		t.FailNow()
	}

}

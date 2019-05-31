package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var tempFile = "./tmp-stdout.data"

func TestWriteResult(t *testing.T) {
	tmpFile, err := os.OpenFile(tempFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name()) // clean up

	if _, err := tmpFile.Seek(0, 0); err != nil {
		panic(err)
	}
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }() // Restore original Stdin

	os.Stdout = tmpFile
	err1 := WriteResult("abc")
	if err1 != nil {
		panic(err1)
	}
	tmpFile.Close()
	data, err := ioutil.ReadFile(tempFile)
	if err != nil {
		panic(err)
	}
	result := string(data)
	fmt.Println(result)
	if result != "abc\n" {
		t.FailNow()
	}
}

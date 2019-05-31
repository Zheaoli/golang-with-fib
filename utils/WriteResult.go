package utils

import (
	"errors"
	"io"
	"os"
)

var TempZoneFull = errors.New("temp zone full")

func WriteResult(result string) error {
	result += "\n"
	number, err := io.WriteString(os.Stdout, result)
	if err != nil {
		return err
	}
	if number != len(result) {
		return TempZoneFull
	}
	return nil
}

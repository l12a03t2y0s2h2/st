package utils

import (
	"errors"
	"os"
)

func CheckFile(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !fileInfo.Mode().IsRegular() {
		return errors.New("%s file isn't a regular file")
	}

	return nil
}

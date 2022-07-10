package tools

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CopyFile(src string, dst string, createDirs bool) error {
	if createDirs {
		os.MkdirAll(filepath.Dir(dst), os.ModeDir|os.ModePerm)
	}

	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

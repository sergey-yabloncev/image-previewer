package cache

import (
	"errors"
	"os"
	"strings"
)

var ErrorClearDir = errors.New("error clear dir")

func RemoveCacheImages(originImagePath, croppedImagePath, name string) error {
	err := clearDir(originImagePath, name)
	if err != nil {
		return ErrorClearDir
	}

	err = clearDir(croppedImagePath, name)
	if err != nil {
		return ErrorClearDir
	}

	return nil
}

func clearDir(dir string, fileName string) error {
	readDirectory, err := os.Open(dir)
	if err != nil {
		return err
	}

	allFiles, err := readDirectory.Readdir(0)
	if err != nil {
		return err
	}

	for _, file := range allFiles {
		if strings.Contains(file.Name(), fileName) {
			err := os.Remove(dir + "/" + file.Name())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

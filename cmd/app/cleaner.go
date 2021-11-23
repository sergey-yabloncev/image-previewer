package main

import (
	"log"
	"os"
)

// Clear paths with cache files before start.
func init() {
	err := clearDir(originImagePath)
	if err != nil {
		log.Fatalf("error clear dir:")
	}

	err = clearDir(croppedImagePath)
	if err != nil {
		log.Fatalf("error clear dir:")
	}
}

// Clear dir exclude .gitignore file.
func clearDir(dir string) error {
	readDirectory, err := os.Open(dir)
	if err != nil {
		return err
	}

	allFiles, err := readDirectory.Readdir(0)
	if err != nil {
		return err
	}

	for _, file := range allFiles {
		if file.Name() == ".gitignore" {
			continue
		}

		err := os.Remove(dir + "/" + file.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

package cleaner

import (
	"log"
	"os"
	"strings"
)

func RemoveCacheImages(originImagePath, croppedImagePath, name string) {
	//Remove origin.
	err := clearDir(originImagePath, name)
	if err != nil {
		log.Fatalf("error clear dir:")
	}

	//Remove cropped.
	err = clearDir(croppedImagePath, name)
	if err != nil {
		log.Fatalf("error clear dir:")
	}
}

// Clear dir exclude .gitignore file.
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

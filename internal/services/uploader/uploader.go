package uploader

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/sergey-yabloncev/image-previewer/internal/helpers"
)

func UploadImage(url, filename, pathDir string, header http.Header) (string, error) {
	imagePath := path.Join(pathDir, filename+".jpg")

	// If file exists return cached image.
	check, err := helpers.IsExists(imagePath)
	if err != nil {
		return "", err
	}

	if check {
		return imagePath, nil
	}

	// Request.
	client := http.Client{}
	ctx := context.Background()
	cancelCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	req, err := http.NewRequestWithContext(cancelCtx, "GET", "http://"+url, nil)
	if err != nil {
		return "", err
	}

	req.Header = header
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("can't download image")
	}

	// Create file.
	file, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Write image.
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return imagePath, nil
}

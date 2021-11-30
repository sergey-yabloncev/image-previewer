package downloader

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

var (
	ErrorNotJpeg      = errors.New("content type isn't image/jpeg")
	ErrorCantDownload = errors.New("can't download image")
)

func DownloadImage(url, filename, pathDir string, header http.Header) (string, error) {
	imagePath := path.Join(pathDir, filename+".jpg")

	check, err := helpers.IsExists(imagePath)
	if err != nil {
		return "", err
	}

	if check {
		return imagePath, nil
	}

	client := http.Client{}
	ctx := context.Background()
	cancelCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	req, err := http.NewRequestWithContext(cancelCtx, http.MethodGet, "http://"+url, nil)
	if err != nil {
		return "", err
	}

	req.Header = header
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if !checkExtension(response, "image/jpeg") {
		return "", ErrorNotJpeg
	}

	if response.StatusCode != 200 {
		return "", ErrorCantDownload
	}

	file, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return imagePath, nil
}

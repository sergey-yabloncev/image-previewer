package router

import (
	"errors"
	"strconv"
	"strings"
)

type CropRequest struct {
	Type   string
	Width  int
	Height int
	URL    string
}

// NewCropRequest returns new instance of Router or false in can't parse url.
func NewCropRequest(requestURI string) (CropRequest, error) {
	path := strings.Split(strings.TrimPrefix(requestURI, "/"), "/")

	if len(path) < 4 {
		return CropRequest{}, errors.New("not enough parameters")
	}

	with, err := strconv.Atoi(path[1])
	if err != nil {
		return CropRequest{}, err
	}

	height, err := strconv.Atoi(path[2])
	if err != nil {
		return CropRequest{}, err
	}

	url := strings.Join(path[3:], "/")
	url = sanitizeURL(url)

	return CropRequest{
		path[0],
		with,
		height,
		url,
	}, nil
}

// Clear host if request has it.
func sanitizeURL(url string) string {
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http:/")
	url = strings.TrimPrefix(url, "https:/")

	return url
}

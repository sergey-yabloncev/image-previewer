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

	return CropRequest{
		path[0],
		with,
		height,
		strings.Join(path[3:], "/"),
	}, nil
}

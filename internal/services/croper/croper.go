package croper

import (
	"fmt"
	"path"

	"github.com/disintegration/imaging"
	"github.com/sergey-yabloncev/image-previewer/internal/helpers"
	"github.com/sergey-yabloncev/image-previewer/internal/router"
)

func Crop(pathImage, pathDir, filename string, request router.CropRequest) (string, error) {
	method := request.Type
	croppedImage := path.Join(pathDir, fmt.Sprintf("%s_%s_%vx%v.jpg", filename, method, request.Width, request.Height))

	check, err := helpers.IsExists(croppedImage)
	if err != nil {
		return "", err
	}

	if check {
		return croppedImage, nil
	}

	src, err := imaging.Open(pathImage)
	if err != nil {
		return "", err
	}

	switch method {
	case "fill":
		src = imaging.Fill(src, request.Width, request.Height, imaging.Center, imaging.Lanczos)
	case "fit":
		src = imaging.Fit(src, request.Width, request.Height, imaging.Lanczos)
	case "anchor":
		src = imaging.CropAnchor(src, request.Width, request.Height, imaging.Center)
	default:
		src = imaging.Resize(src, request.Width, request.Height, imaging.Lanczos)
	}

	err = imaging.Save(src, croppedImage)
	if err != nil {
		return "", err
	}

	return croppedImage, nil
}

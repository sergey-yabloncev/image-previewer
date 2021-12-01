package croper

import (
	"github.com/disintegration/imaging"
	"github.com/sergey-yabloncev/image-previewer/internal/router"
)

func Crop(pathImage, filename string, request router.CropRequest) error {
	method := request.Type

	src, err := imaging.Open(pathImage)
	if err != nil {
		return err
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

	err = imaging.Save(src, filename)
	if err != nil {
		return err
	}

	return nil
}

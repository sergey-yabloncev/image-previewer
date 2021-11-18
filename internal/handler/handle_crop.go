package handler

import (
	"net/http"
	"strings"

	"github.com/sergey-yabloncev/image-previewer/internal/router"
	"github.com/sergey-yabloncev/image-previewer/internal/services/croper"
	"github.com/sergey-yabloncev/image-previewer/internal/services/uploader"
)

type CropHandler struct {
	originImagePath  string
	croppedImagePath string
}

func NewCropHandler(originImagePath, croppedImagePath string) CropHandler {
	return CropHandler{
		originImagePath,
		croppedImagePath,
	}
}

// Main resolver function.
func (h CropHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !checkExtension(r.RequestURI) {
		router.HTTPBadRequest(w, "Can't find image", nil)
		return
	}

	request, err := router.NewCropRequest(r.RequestURI)
	if err != nil {
		router.HTTPBadRequest(w, "Can't parse parameters", err)
		return
	}

	srcOriginImage, err := uploader.UploadImage(request.URL, h.originImagePath, r.Header)
	if err != nil {
		router.HTTPInternalServerError(w, "Can't upload image", err)
		return
	}

	outImage, err := croper.Crop(srcOriginImage, h.croppedImagePath, request)
	if err != nil {
		router.HTTPInternalServerError(w, "Can't crop image", err)
		return
	}

	http.ServeFile(w, r, outImage)
}

// Find jpg extension in request RequestURI.
func checkExtension(url string) bool {
	return strings.Contains(url, ".jpg") || strings.Contains(url, ".jpeg")
}

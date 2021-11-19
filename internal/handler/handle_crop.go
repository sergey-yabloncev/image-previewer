package handler

import (
	"github.com/sergey-yabloncev/image-previewer/internal/helpers"
	"github.com/sergey-yabloncev/image-previewer/internal/services/cache"
	"net/http"
	"strings"

	"github.com/sergey-yabloncev/image-previewer/internal/router"
	"github.com/sergey-yabloncev/image-previewer/internal/services/cleaner"
	"github.com/sergey-yabloncev/image-previewer/internal/services/croper"
	"github.com/sergey-yabloncev/image-previewer/internal/services/uploader"
)

type CropHandler struct {
	originImagePath  string
	croppedImagePath string
	cache            cache.Cache
}

func NewCropHandler(originImagePath, croppedImagePath string, cache cache.Cache) CropHandler {
	return CropHandler{
		originImagePath,
		croppedImagePath,
		cache,
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
	// Generate new uniq name.
	fileName := helpers.Md5String(request.URL)
	// Upload image.
	srcOriginImage, err := uploader.UploadImage(request.URL, fileName, h.originImagePath, r.Header)
	if err != nil {
		router.HTTPInternalServerError(w, "Can't upload image", err)
		return
	}
	// Generate cropped image.
	outImage, err := croper.Crop(srcOriginImage, fileName, h.croppedImagePath, request)
	if err != nil {
		router.HTTPInternalServerError(w, "Can't crop image", err)
		return
	}

	// Set image to cache.
	_, removedImage := h.cache.Set(cache.Key(fileName), "")
	//If we have removed item, we're removing from disk with cropped images
	if removedImage != "" {
		cleaner.RemoveCacheImages(h.originImagePath, h.croppedImagePath, fileName)
	}

	http.ServeFile(w, r, outImage)
}

// Find jpg extension in request RequestURI.
func checkExtension(url string) bool {
	return strings.Contains(url, ".jpg") || strings.Contains(url, ".jpeg")
}

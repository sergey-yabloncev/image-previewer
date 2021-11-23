package handler

import (
	"github.com/sergey-yabloncev/image-previewer/internal/services/downloader"
	"log"
	"net/http"
	"strings"

	"github.com/sergey-yabloncev/image-previewer/internal/helpers"
	"github.com/sergey-yabloncev/image-previewer/internal/router"
	"github.com/sergey-yabloncev/image-previewer/internal/services/cache"
	"github.com/sergey-yabloncev/image-previewer/internal/services/cleaner"
	"github.com/sergey-yabloncev/image-previewer/internal/services/croper"
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

	fileName := helpers.Md5String(request.URL)
	srcOriginImage, err := downloader.DownloadImage(request.URL, fileName, h.originImagePath, r.Header)
	if err != nil {
		router.HTTPInternalServerError(w, "Can't upload image", err)
		return
	}

	outImage, err := croper.Crop(srcOriginImage, h.croppedImagePath, fileName, request)
	if err != nil {
		router.HTTPInternalServerError(w, "Can't crop image", err)
		return
	}

	_, removedImage := h.cache.Set(cache.Key(fileName), "")
	if removedImage != "" {
		log.Println("Images was removed:", removedImage)
		cleaner.RemoveCacheImages(h.originImagePath, h.croppedImagePath, string(removedImage))
	}

	http.ServeFile(w, r, outImage)
}

// Find jpg extension in request RequestURI.
func checkExtension(url string) bool {
	return strings.Contains(url, ".jpg") || strings.Contains(url, ".jpeg")
}

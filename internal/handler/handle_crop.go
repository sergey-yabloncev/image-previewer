package handler

import (
	"net/http"

	"github.com/sergey-yabloncev/image-previewer/internal/helpers"
	"github.com/sergey-yabloncev/image-previewer/internal/router"
	"github.com/sergey-yabloncev/image-previewer/internal/services/cache"
	"github.com/sergey-yabloncev/image-previewer/internal/services/croper"
	"github.com/sergey-yabloncev/image-previewer/internal/services/downloader"
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
	request, err := router.NewCropRequest(r.RequestURI)
	if err != nil {
		router.HTTPBadRequest(w, "Can't parse parameters", err)
		return
	}

	fileName := helpers.Md5String(request.URL)
	srcOriginImage, err := downloader.DownloadImage(request.URL, fileName, h.originImagePath, r.Header)
	if err != nil {
		router.HTTPBadRequest(w, "Can't upload image", err)
		return
	}

	h.cache.Set(cache.Key(fileName), "")

	outImage, err := croper.Crop(srcOriginImage, h.croppedImagePath, fileName, request)
	if err != nil {
		router.HTTPInternalServerError(w, "Can't crop image", err)
		return
	}

	http.ServeFile(w, r, outImage)
}

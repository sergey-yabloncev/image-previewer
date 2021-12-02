package handler

import (
	"fmt"
	"net/http"
	"path"
	"strings"

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
	request, err := router.NewCropRequest(r.URL.Path)
	if err != nil {
		router.HTTPBadRequest(w, "Can't parse parameters", err)
		return
	}

	fileName := helpers.Hash(request.URL)
	originImagePath := path.Join(h.originImagePath, fileName+".jpg")
	if _, ok := h.cache.Get(cache.Key(fileName)); !ok {
		err := downloader.DownloadImage(request.URL, originImagePath, r.Header)
		if err != nil {
			router.HTTPBadRequest(w, "Can't upload image", err)
			return
		}

		h.cache.Set(cache.Key(fileName), "")
	}

	croppedImage := path.Join(
		h.croppedImagePath,
		fmt.Sprintf("%s_%s_%vx%v.jpg", fileName, request.Type, request.Width, request.Height),
	)
	cacheValue, ok := h.cache.Get(cache.Key(fileName))
	if ok {
		if strings.Contains(fmt.Sprint(cacheValue), croppedImage) {
			http.ServeFile(w, r, croppedImage)
			return
		}
	}

	err = croper.Crop(originImagePath, croppedImage, request)
	if err != nil {
		router.HTTPInternalServerError(w, "Can't crop image", err)
		return
	}

	h.cache.Set(cache.Key(fileName), croppedImage+"|"+fmt.Sprint(cacheValue))

	http.ServeFile(w, r, croppedImage)
}

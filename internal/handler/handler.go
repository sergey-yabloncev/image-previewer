package handler

import (
	"net/http"
	"strings"
)

type Handler struct {
	cropHandler CropHandler
	docHandler  DocHandler
}

func New(cropHandler CropHandler, docHandler DocHandler) Handler {
	return Handler{cropHandler, docHandler}
}

// Main resolver function.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rootPath := getRootPath(r.RequestURI)
	switch rootPath {
	case "fill", "fit", "anchor", "default":
		h.cropHandler.ServeHTTP(w, r)
	case "":
		h.docHandler.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

// Get first part path.
func getRootPath(requestURI string) string {
	return strings.Split(strings.TrimPrefix(requestURI, "/"), "/")[0]
}

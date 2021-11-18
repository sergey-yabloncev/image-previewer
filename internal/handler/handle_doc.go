package handler

import (
	"github.com/sergey-yabloncev/image-previewer/internal/router"
	"html/template"
	"net/http"
)

type DocHandler struct {
	originImagePath  string
	croppedImagePath string
}

func NewDocHandler() DocHandler {
	return DocHandler{}
}

// Main resolver function.
func (h DocHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		router.HTTPBadRequest(w, "Internal Server Error", err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		router.HTTPBadRequest(w, "Internal Server Error", err)
		return
	}
}

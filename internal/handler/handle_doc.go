package handler

import (
	"html/template"
	"net/http"

	"github.com/sergey-yabloncev/image-previewer/internal/router"
)

type DocHandler struct{}

func NewDocHandler() DocHandler {
	return DocHandler{}
}

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

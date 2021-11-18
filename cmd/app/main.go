package main

import (
	"log"
	"net/http"
	"path"

	"github.com/sergey-yabloncev/image-previewer/internal/handler"
	"github.com/sergey-yabloncev/image-previewer/internal/router"
)

var (
	// Dirs for storage cache images.
	originImagePath  = path.Join("storage", "public", "origin")
	croppedImagePath = path.Join("storage", "public", "cropped")
)

func main() {
	server := http.NewServeMux()
	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./storage/public"))))

	rootHandler := handler.New(
		handler.NewCropHandler(originImagePath, croppedImagePath),
		handler.NewDocHandler(),
	)

	server.HandleFunc("/", rootHandler.ServeHTTP)

	log.Println("start server")
	log.Fatal(http.ListenAndServe(":8080", router.LoggerMiddleware(server)))
}

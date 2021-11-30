package main

import (
	"log"
	"net/http"

	"github.com/sergey-yabloncev/image-previewer/internal/handler"
	"github.com/sergey-yabloncev/image-previewer/internal/router"
	"github.com/sergey-yabloncev/image-previewer/internal/services/cache"
)

func Server(config Config) {
	server := http.NewServeMux()
	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./storage/public"))))

	rootHandler := handler.New(
		handler.NewCropHandler(
			originImagePath,
			croppedImagePath,
			cache.NewCache(
				config.Cache.Capacity,
				originImagePath,
				croppedImagePath,
				true,
			)),
		handler.NewDocHandler(),
	)

	server.HandleFunc("/", rootHandler.ServeHTTP)

	log.Println("start server to http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router.LoggerMiddleware(server)))
}

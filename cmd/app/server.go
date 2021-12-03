package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sergey-yabloncev/image-previewer/internal/handler"
	"github.com/sergey-yabloncev/image-previewer/internal/router"
	"github.com/sergey-yabloncev/image-previewer/internal/services/cache"
)

func Server(config Config) {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./storage/public"))))

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

	mux.HandleFunc("/", rootHandler.ServeHTTP)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router.LoggerMiddleware(mux),
	}

	log.Println("start server to http://localhost:8080/")
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}

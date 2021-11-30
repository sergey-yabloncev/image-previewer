package main

import (
	"log"
	"path"
)

var (
	originImagePath  = path.Join("storage", "public", "origin")
	croppedImagePath = path.Join("storage", "public", "cropped")
)

func main() {
	config, err := readConfig("./configs/config.toml")
	if err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	Server(config)
}

package main

import "github.com/BurntSushi/toml"

func readConfig(path string) (c Config, err error) {
	_, err = toml.DecodeFile(path, &c)
	return
}

type Config struct {
	Cache CacheConfig
}

type CacheConfig struct {
	Capacity int
}

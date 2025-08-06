package main

import (
	"flag"
	"log"

	"github.com/!burnt!sushi/toml"
	"github.com/iktkhor/http-rest-api.git/internal/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	s := apiserver.New(config)
	if err = s.Start(); err != nil {
		log.Fatal(err)
	}
}

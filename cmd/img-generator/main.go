package main

import (
	"flag"
	"log"

	"github.com/D-Zane/img-generator/configs"
	"github.com/D-Zane/img-generator/internal/server"
)

var confPath = flag.String("conf-path", "./configs/.env", "Path to config env")

func main() {
	conf, err := configs.New(*confPath)
	if err != nil {
		log.Fatalln(err)
	}
	server.Run(conf)
}

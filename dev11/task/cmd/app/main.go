package main

import (
	"dev11/internal/config"
	"dev11/internal/hash"
	"dev11/internal/server"
	"flag"
	"log"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", ":8080", "use for setting port, by default its :8080")
}

func main() {
	flag.Parse()

	conf := config.NewConfig()

	hash, err := hash.NewHash()
	if err != nil {
		log.Fatal(err)
	}

	app := server.NewApp(conf, hash)

	if err = app.Start(); err != nil {
		log.Fatal(err)
	}
}

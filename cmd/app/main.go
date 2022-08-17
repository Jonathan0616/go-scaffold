package main

import (
	"log"

	"go-scaffold/config"
	"go-scaffold/internal/app"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}

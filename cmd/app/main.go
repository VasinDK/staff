package main

import (
	"log"
	"this_module/config"
	"this_module/internal/app"
)

func main() {
	// Получаем конфиги
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}

package main

import (
	"log"
	"this_module/config"
	"this_module/internal/app"
)

func main() {
	// Конфигурация приложения
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Запуск приложения
	app.Run(cfg)
}

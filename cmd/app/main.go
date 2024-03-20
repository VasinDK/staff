package main

import (
	"log"
	"this_module/config"
	_ "this_module/docs"
	"this_module/internal/app"
)

//	@title Документация для проекта staff
//	@version 1.0
//	@description Web-Сервис сотрудников. Сервис добавляет сотрудников,
//  @description удаляет по id, выводит список для указанной компании,список для отдела. Изменяет сотрудников по id

// 	@host localhost:8080
// 	@BasePath /

func main() {
	// Конфигурация приложения
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Запуск приложения
	app.Run(cfg)
}

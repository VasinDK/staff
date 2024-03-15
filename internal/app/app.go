// Основной пакет приложения.
// В основе лежит чистая архитектура. Слои:
//   - Entity. Все основные сущности
//   - uscase. Пользовательские кейсы
//   - controller. Контроллер. Сейчас только HTTP, но можно добавить еще
//
// Тут создаются основные зависимости приложения:
//   - Логер
//   - Пул соединения с бд
//   - uscase использует пакет с утилитами и репозиторий (с бд postgresql в основе)
//   - Роутер создан на основе пакета Chi
//
// httpserver запускается в отдельной горутине. Доступно прослушивание os.Interrupt, syscall.SIGTERM
// и остановка сервера с помощью Shutdown
package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"this_module/config"
	v1 "this_module/internal/controller/http/v1"
	"this_module/internal/infrastructure/repository"
	"this_module/internal/pkg/utils"
	"this_module/internal/usecase"
	"this_module/pkg/httpserver"
	"this_module/pkg/logger"
	"this_module/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	// Репозиторий
	db, err := postgres.New(cfg.Storage)
	if err != nil {
		l.Error("postgres.New", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Use case
	staffUseCase := usecase.New(
		utils.New(),
		repository.New(db),
	)

	// router
	router, err := v1.NewRouter(l, staffUseCase)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}

	// HTTP сервер
	s := httpserver.New(router, cfg)

	// Ожидание сигнала завершения работы от ОС
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	select {
	case err = <-s.Notify:
		l.Error(err.Error())
	case <-ctx.Done():
		l.Info("signal Interrupt")
	}

	// Shutdown
	s.Stop()

	l.Info("Server stoped")
}

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

	db, err := postgres.New(cfg.Storage)
	if err != nil {
		l.Error("postgres.New", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	staffUseCase := usecase.New(
		utils.New(),
		repository.New(db),
	)

	router, err := v1.NewRouter(l, staffUseCase)
	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}

	s := httpserver.New(router, cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	select {
	case err = <-s.Notify:
		l.Error(err.Error())
	case <-ctx.Done():
		l.Info("signal Interrupt")
	}

	s.Stop()

	l.Info("Server stoped")
}

package postgres

import (
	"fmt"
	"this_module/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func New(cfg config.Storage) (*sqlx.DB, error) {
	stringConnect := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Pass, cfg.Addr, cfg.Port, cfg.Name,
	)

	conn, err := sqlx.Connect("postgres", stringConnect)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx connect")
	}

	err = conn.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "sqlx ping")
	}

	return conn, nil
}

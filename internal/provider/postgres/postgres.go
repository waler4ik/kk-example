// Code generated by kk; BUT FEEL FREE TO EDIT.

package postgres

import (
	"database/sql"
	"fmt"

	"github.com/caarlos0/env/v10"
	_ "github.com/lib/pq"
	"github.com/waler4ik/kk-example/internal/secretmanager"
)

type config struct {
	Address string `env:"POSTGRES_ADDRESS" envDefault:"db:5432"`
	User    string `env:"POSTGRES_USER" envDefault:"postgres"`
	DBName  string `env:"POSTGRES_DBNAME" envDefault:"postgres"`
}

type Postgres struct {
	log func(string, ...any)
	db  *sql.DB
}

func New(logger func(string, ...any), sm secretmanager.SecretManager) (*Postgres, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}
	pwd, err := sm.Secret("POSTGRES_PWD")
	if err != nil {
		return nil, fmt.Errorf("secret: %w", err)
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, pwd, cfg.Address, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return &Postgres{
		log: logger,
		db:  db,
	}, nil
}

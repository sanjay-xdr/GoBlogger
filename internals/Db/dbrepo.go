package Db

import (
	"database/sql"

	"github.com/sanjay-xdr/goblogger/internals/config"
)

type PostgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) *PostgresDBRepo {
	return &PostgresDBRepo{
		App: a,
		DB:  conn,
	}
}

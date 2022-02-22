package dbrepo

import (
	"database/sql"

	"github.com/SyedAliHamad/internproject/internal/repository"
	"github.com/SyedAliHamad/internproject/pkg/config"
)

type postgresDBrepo struct{
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB,a *config.AppConfig)repository.DatabaseRepo{

	return &postgresDBrepo{
		App:a,
		DB:conn,
	}
}
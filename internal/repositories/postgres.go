package repositories

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/web1/internal/models"
)

const (
	usersTable = "users"
)

func NewPostgresDB(cfg *models.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.ConnectDb)
	if err != nil {
		return nil, err
	}
	return db, nil
}

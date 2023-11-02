package repositories

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/web1/internal/models"
)

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (*models.User, error)
}

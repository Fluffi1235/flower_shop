package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/web1/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, login, password) values ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Email, user.Login, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *AuthPostgres) GetUser(login, password string) (*models.User, error) {
	user := new(models.User)
	query := fmt.Sprintf("SELECT id FROM %s WHERE login = $1 AND password = $2", usersTable)

	err := s.db.Get(user, query, login, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

package services

import (
	"gitlab.com/web1/internal/models"
	"gitlab.com/web1/internal/repositories"
)

type Service struct {
	Authorization
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

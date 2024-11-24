package service

import (
	"music/pkg/repository"
)

type Authorization interface {
	SingUp(username string, password string) error
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repo),
	}
}

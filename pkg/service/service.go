package service

import "music/pkg/repository"

type Authorization interface {
	SingUp(username string, password string) error
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repo),
	}
}

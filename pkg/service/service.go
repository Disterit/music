package service

import "music/pkg/repository"

type Auth interface {

}

type Service struct {
	Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{

	}
}

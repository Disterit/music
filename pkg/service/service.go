package service

import (
	"music"
	"music/pkg/repository"
)

type Authorization interface {
	SingUp(username string, password string) error
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Song interface {
	CreateSong(song music.Song) error
}

type Service struct {
	Authorization
	Song
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repo),
		Song:          NewSongService(repo),
	}
}

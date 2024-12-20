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

type Album interface {
	CreateAlbum(album music.Album) (int, error)
	GetAlbums(artistId int) ([]music.Album, error)
	GetAlbum(artistId, albumId int) (music.Album, error)
	UpdateAlbum(album music.Album) error
	DeleteAlbum(artistId, albumId int) error
}

type Song interface {
	CreateSong(song music.Song, artistId int) (int, error)
}

type Service struct {
	Authorization
	Album
	Song
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repo),
		Album:         NewAlbumService(repo),
		Song:          NewSongService(repo),
	}
}

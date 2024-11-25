package repository

import (
	"github.com/jmoiron/sqlx"
	"music"
)

type Authorization interface {
	CreateArtist(username string, password string) error
	GetArtist(username, password string) (music.Artist, error)
}

type Album interface {
	CreateAlbum(album music.Album) (int, error)
}

type Song interface {
	CreateSong(song music.Song) error
}

type Repository struct {
	Authorization
	Song
	Album
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Album:         NewAlbumRepository(db),
		Song:          NewSongRepository(db),
	}
}

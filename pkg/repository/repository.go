package repository

import (
	"github.com/jmoiron/sqlx"
	"music"
)

type Authorization interface {
	CreateArtist(username string, password string) error
	GetArtist(username, password string) (music.Artist, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}

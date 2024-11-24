package repository

import (
	"github.com/jmoiron/sqlx"
	"music"
)

type SongRepository struct {
	db *sqlx.DB
}

func NewSongRepository(db *sqlx.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) CreateSong(song music.Song) error {

	return nil
}

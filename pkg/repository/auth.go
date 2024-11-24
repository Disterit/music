package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"music"
	"music/logger"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateArtist(username string, password string) error {
	query := fmt.Sprintf("INSERT INTO %s (artist_name, password_hash) VALUES ($1, $2);", artistTable)
	_, err := r.db.Exec(query, username, password)
	if err != nil {
		logger.Log.Error("error to create artist", err.Error())
		return err
	}
	return nil
}

func (r *AuthRepository) GetArtist(username, password string) (music.Artist, error) {
	var artist music.Artist
	query := fmt.Sprintf("SELECT id FROM %s WHERE artist_name = $1 AND password_hash = 2$", artistTable)

	err := r.db.Get(&artist, query, username, password)
	if err != nil {
		logger.Log.Error("error to get artist", err.Error())
		return artist, err
	}

	return artist, nil
}

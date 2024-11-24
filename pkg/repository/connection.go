package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"music/logger"

	_ "github.com/lib/pq"
)

const (
	artistTable         = "artist"
	genreTable          = "genre"
	albumTable          = "album"
	songsTable          = "songs"
	playlistTable       = "playlist"
	playlist_songsTable = "playlist_songs"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func Connection(cfg Config) *sqlx.DB {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.SSLMode))

	if err != nil {
		logger.Log.Error("error to connect to database", err.Error())
		return nil
	}

	return db
}

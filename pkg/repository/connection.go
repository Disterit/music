package repository

import (
	"database/sql"
	"fmt"
	"music/logger"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func Connection(cfg Config) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.SSLMode))

	if err != nil {
		logger.Log.Error("error to connect to database", err.Error())
		return nil
	}

	return db
}

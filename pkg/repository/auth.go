package repository

import (
	"database/sql"
	"fmt"
	"music/logger"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) SingUp(username string, password string) error {
	query := fmt.Sprintf("INSERT INTO %s (artist_name, password_hash) VALUES ($1, $2);", artistTable)
	_, err := r.db.Exec(query, username, password)
	if err != nil {
		logger.Log.Error("error to create artist", err.Error())
		return err
	}
	return nil
}

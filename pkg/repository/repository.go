package repository

import "database/sql"

type Authorization interface {
	SingUp(username string, password string) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}

package repository

import "database/sql"

type Auth interface {

}

type Repository struct {
	Auth
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{

	}
}
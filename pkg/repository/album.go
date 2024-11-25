package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"music"
)

type AlbumRepository struct {
	db *sqlx.DB
}

func NewAlbumRepository(db *sqlx.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (r *AlbumRepository) CreateAlbum(album music.Album) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (title_album, id_artist) VALUES ($1, $2) RETURNING id`, albumTable)
	row := r.db.QueryRow(query, album.Title, album.IdArtist)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

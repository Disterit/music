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

func (r *AlbumRepository) GetAlbums(artistId int) ([]music.Album, error) {
	var albums []music.Album

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id_artist = $1`, albumTable)
	rows, err := r.db.Query(query, artistId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album music.Album
		if err := rows.Scan(&album.Id, &album.Title, &album.IdArtist); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumRepository) GetAlbum(artistId, albumId int) (music.Album, error) {
	var album music.Album

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1 AND id_artist = $2`, albumTable)
	row := r.db.QueryRow(query, albumId, artistId)

	if err := row.Scan(&album.Id, &album.Title, &album.IdArtist); err != nil {
		return album, err
	}

	return album, nil
}

func (r *AlbumRepository) UpdateAlbum(album music.Album) error {
	query := fmt.Sprintf(`UPDATE %s SET title_album = $3 WHERE id = $1 AND id_artist = $2 `, albumTable)

	_, err := r.db.Exec(query, album.Id, album.IdArtist, album.Title)

	return err
}

func (r *AlbumRepository) DeleteAlbum(artistId, albumId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 AND id_artist = $2`, albumTable)
	_, err := r.db.Exec(query, albumId, artistId)
	return err
}

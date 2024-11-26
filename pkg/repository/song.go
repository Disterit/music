package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"music"
)

type SongRepository struct {
	db *sqlx.DB
}

func NewSongRepository(db *sqlx.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) CreateSong(song music.Song, artistId int) (int, error) {
	var id int

	query := fmt.Sprintf(`SELECT id FROM %s WHERE id = $1 AND id_artist = $2`, albumTable)
	row := r.db.QueryRow(query, song.AlbumId, artistId)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Not your album: %w", err)
	}

	query = fmt.Sprintf(`INSERT INTO %s (title_song, text_song, id_genre, id_album) VALUES ($1, $2, $3, $4) RETURNING id`, songsTable)
	row = r.db.QueryRow(query, song.TitleSong, song.TextSong, song.GenreID, song.AlbumId)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("Not your song: %w", err)
	}

	return id, nil
}

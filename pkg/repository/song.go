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

	query := fmt.Sprintf(`
		WITH valid_album AS (
			SELECT id FROM %s WHERE id = $1 AND id_artist = $2
		)
		INSERT INTO %s (title_song, text_song, id_genre, id_album)
		SELECT $3, $4, $5, id 
		FROM valid_album
		RETURNING id
	`, albumTable, songsTable)

	row := r.db.QueryRow(query, song.AlbumId, artistId, song.TitleSong, song.TextSong, song.GenreID)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("Failed to create song: %w", err)
	}

	return id, nil
}

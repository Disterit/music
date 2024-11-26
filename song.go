package music

type Song struct {
	Id        int    `json:"id"`
	TitleSong string `json:"title_song" binding:"required"`
	TextSong  string `json:"text_song" binding:"required"`
	GenreID   int    `json:"genreID" binding:"required"`
	AlbumId   int    `json:"albumID" binding:"required"`
}

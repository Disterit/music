package music

type Song struct {
	Id        int    `json:"id"`
	TitleSong string `json:"title_song" binding:"required"`
	TextSong  string `json:"text_song" binding:"required"`
	Genre     string `json:"genre" binding:"required"`
}

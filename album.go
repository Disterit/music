package music

type Album struct {
	Id       string `json:"id"`
	Title    string `json:"title" binding:"required"`
	IdArtist int    `json:"id_artist"`
}

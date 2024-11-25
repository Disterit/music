package music

type Album struct {
	Id       int    `json:"id"`
	Title    string `json:"title" binding:"required"`
	IdArtist int    `json:"id_artist"`
}

package music

type Artist struct {
	Id         int    `json:"id"`
	ArtistName string `json:"name"`
	Password   string `json:"password"`
}

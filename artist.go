package music

type Artist struct {
	Id         int    `json:"id"`
	ArtistName string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"music"
	"music/logger"
	"net/http"
)

func (h *Handler) AddSong(c *gin.Context) {
	var inputSong music.Song
	artistId, err := getArtistId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to read artist id")
		return
	}

	if err := c.ShouldBindJSON(&inputSong); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to add song")

	}

}

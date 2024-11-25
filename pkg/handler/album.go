package handler

import (
	"github.com/gin-gonic/gin"
	"music"
	"music/logger"
	"net/http"
)

func (h *Handler) CreateAlbum(c *gin.Context) {
	artistId, err := getArtistId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get artist ID")
		return
	}

	var inputAlbum music.Album
	inputAlbum.IdArtist = artistId
	if err = c.ShouldBindJSON(&inputAlbum); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to bind input data")
		return
	}

	id, err := h.service.Album.CreateAlbum(inputAlbum)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to create album")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

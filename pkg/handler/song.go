package handler

import (
	"github.com/gin-gonic/gin"
	"music"
	"music/logger"
	"net/http"
)

func (h *Handler) CreateSong(c *gin.Context) {
	var inputSong music.Song
	artistID, err := getArtistId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get artist")
		return
	}

	if err = c.ShouldBindJSON(&inputSong); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to add song", err.Error())

	}

	id, err := h.service.Song.CreateSong(inputSong, artistID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to add song", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) GetSong(c *gin.Context) {}

func (h *Handler) GetAllSongArtist(c *gin.Context) {}

func (h *Handler) GetAllSong(c *gin.Context) {}

func (h *Handler) UpdateSong(c *gin.Context) {}

func (h *Handler) DeleteSong(c *gin.Context) {}

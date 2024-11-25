package handler

import (
	"github.com/gin-gonic/gin"
	"music"
	"music/logger"
	"net/http"
	"strconv"
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

func (h *Handler) GetAlbums(c *gin.Context) {
	artistId, err := getArtistId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get artist ID")
		return
	}

	albums, err := h.service.Album.GetAlbums(artistId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get albums")
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (h *Handler) GetAlbum(c *gin.Context) {
	artistId, err := getArtistId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get artist ID")
		return
	}

	albumId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get id of album")
		return
	}

	album, err := h.service.Album.GetAlbum(artistId, albumId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get album", err.Error())
		return
	}

	c.JSON(http.StatusOK, album)
}

func (h *Handler) UpdateAlbums(c *gin.Context) {
	artistId, err := getArtistId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get artist ID")
		return
	}

	albumId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get id of album")
		return
	}

	var inputAlbum music.Album
	inputAlbum.IdArtist = artistId
	inputAlbum.Id = albumId

	if err = c.ShouldBindJSON(&inputAlbum); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to bind input data")
		return
	}

	err = h.service.Album.UpdateAlbum(inputAlbum)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to update album", err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) DeleteAlbums(c *gin.Context) {
	artistId, err := getArtistId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get artist ID")
		return
	}

	albumId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to get id of album")
		return
	}

	err = h.service.Album.DeleteAlbum(artistId, albumId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error("error to delete album", err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

package handler

import (
	"github.com/gin-gonic/gin"
	"music"
	"music/logger"
	"net/http"
)

func (h *Handler) SingUp(c *gin.Context) {
	var input music.Artist

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to read input signUp")
		return
	}

	err := h.service.SingUp(input.ArtistName, input.PasswordHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) SingIn(c *gin.Context) {

}

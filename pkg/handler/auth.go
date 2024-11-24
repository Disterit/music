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

	err := h.service.SingUp(input.ArtistName, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type singInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SingIn(c *gin.Context) {
	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logger.Log.Error("error to read input signIn")
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

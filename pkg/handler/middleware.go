package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"music/logger"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	artistCtx           = "ArtistId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "No authorization header")
		logger.Log.Error("No authorization header found")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		logger.Log.Error("Invalid authorization header")
		return
	}

	artistId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		logger.Log.Error("Invalid authorization header")
		return
	}

	c.Set(artistCtx, artistId)
}

func getArtistId(c *gin.Context) (int, error) {
	id, ok := c.Get(artistCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "No artist found")
		logger.Log.Error("No artist found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id is of invalid type")
		logger.Log.Error("user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

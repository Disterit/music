package handler

import "github.com/gin-gonic/gin"

type errorResponse struct {
	message string `json:"message" :"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, status int, msg string) {
	c.AbortWithStatusJSON(status, errorResponse{msg})
}

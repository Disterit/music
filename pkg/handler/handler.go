package handler

import (
	"github.com/gin-gonic/gin"
	"music/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.SingUp)
		auth.POST("/sing-in", h.SingIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		song := api.Group("/song")
		{
			song.POST("/")
			song.GET("/")
		}
	}

	return router
}

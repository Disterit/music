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
			song.POST("/", h.CreateSong)
			song.GET("/", h.GetAllSong)
			song.GET("/:id", h.GetSong)
			song.PUT("/:id", h.UpdateSong)
			song.DELETE("/:id", h.DeleteSong)
			song.GET("/artist/:id", h.GetAllSongArtist)

		}

		album := api.Group("/album")
		{
			album.POST("/", h.CreateAlbum)
			album.GET("/", h.GetAlbums)
			album.GET("/:id", h.GetAlbum)
			album.DELETE("/:id", h.DeleteAlbums)
			album.PUT("/:id", h.UpdateAlbums)
		}
	}

	return router
}

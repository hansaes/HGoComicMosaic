package http

import (
	"HGoComicMosaic/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(api *gin.RouterGroup, h *handler.AuthHandler) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.GET("/me", h.Me)
	}
}

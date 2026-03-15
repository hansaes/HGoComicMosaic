package http

import (
	"HGoComicMosaic/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterResourceRoutes(api *gin.RouterGroup, h *handler.ResourceHandler) {
	resources := api.Group("/resources")
	{
		resources.GET("", h.List)
		resources.GET("/:id", h.Detail)
		resources.POST("", h.Create)
	}
}

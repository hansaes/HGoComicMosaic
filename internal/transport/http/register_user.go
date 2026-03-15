package http

import (
	"HGoComicMosaic/internal/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup, h *handler.UserHandler) {
	admin := api.Group("/admin/users")
	{
		admin.POST("", h.Create)
	}
}

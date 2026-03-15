package http

import (
	"HGoComicMosaic/internal/transport/http/handler"
	"HGoComicMosaic/internal/transport/http/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth     *handler.AuthHandler
	User     *handler.UserHandler
	Resource *handler.ResourceHandler
}

func NewRouter(h Handlers) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(middleware.CROS())

	r.GET("/healthz", func(c *gin.Context) {
		requestID, _ := c.Get(middleware.RequestIDKey)
		c.JSON(200, gin.H{
			"status":     "ok",
			"request_id": requestID,
		})
	})

	apiV1 := r.Group("/api/v1")

	RegisterAuthRoutes(apiV1, h.Auth)
	RegisterUserRoutes(apiV1, h.User)
	RegisterResourceRoutes(apiV1, h.Resource)

	return r

}

package http

import (
	"HGoComicMosaic/internal/transport/http/handler"
	"HGoComicMosaic/internal/transport/http/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
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

	authHandler := handler.NewAuthHandler()
	userHandler := handler.NewUserHandler()
	resourceHandler := handler.NewResourceHandler()

	RegisterAuthRoutes(apiV1, authHandler)
	RegisterUserRoutes(apiV1, userHandler)
	RegisterResourceRoutes(apiV1, resourceHandler)

	return r

}

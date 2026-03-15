package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceHandler struct {
}

func NewResourceHandler() *ResourceHandler {
	return &ResourceHandler{}
}

func (h *ResourceHandler) List(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "list not implemented"})
}

func (h *ResourceHandler) Detail(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "detail not implemented"})
}

func (h *ResourceHandler) Create(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "create not implemented"})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "login not implemented"})
}

func (h *AuthHandler) Me(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "me not implemented"})
}

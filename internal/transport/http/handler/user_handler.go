package handler

import (
	"HGoComicMosaic/internal/repository"
	"HGoComicMosaic/internal/service"
	"HGoComicMosaic/internal/transport/http/dto"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), req.Username, req.Password, req.IsAdmin)
	if err != nil {
		if errors.Is(err, repository.ErrUsernameExists) {
			c.JSON(http.StatusConflict, gin.H{
				"message": "username already exists",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "create user failed",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	})
}

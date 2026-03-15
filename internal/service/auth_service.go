package service

import (
	platformauth "HGoComicMosaic/internal/platform/auth"
	"HGoComicMosaic/internal/repository"
	"context"
	"errors"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	userRepo     repository.UserRepository
	tokenService *platformauth.TokenService
}

func NewAuthService(userRepo repository.UserRepository, tokenService *platformauth.TokenService) *AuthService {
	return &AuthService{userRepo: userRepo, tokenService: tokenService}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, int64, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, repository.ErrUsernameExists) {
			return "", 0, ErrInvalidCredentials
		}
		return "", 0, err
	}

	if !platformauth.VerifyPassword(user.HashedPassword, password) {
		return "", 0, ErrInvalidCredentials
	}

	return s.tokenService.GenerateToken(user.ID, username, user.IsAdmin)
}

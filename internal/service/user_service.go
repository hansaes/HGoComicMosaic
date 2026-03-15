package service

import (
	"HGoComicMosaic/internal/domain"
	platformauth "HGoComicMosaic/internal/platform/auth"
	"HGoComicMosaic/internal/repository"
	"context"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, username, password string, isAdmin bool) (*domain.User, error) {
	hashedPassword, err := platformauth.HashPassword(password)

	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Username:       username,
		HashedPassword: hashedPassword,
		IsAdmin:        isAdmin,
	}

	if err = s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

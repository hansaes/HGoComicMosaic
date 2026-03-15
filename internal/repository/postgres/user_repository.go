package postgres

import (
	"HGoComicMosaic/internal/domain"
	"HGoComicMosaic/internal/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type userModel struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Username       string    `gorm:"column:username;size:64;not null;uniqueIndex"`
	HashedPassword string    `gorm:"column:hashed_password;type:text;not null"`
	IsAdmin        bool      `gorm:"column:is_admin;not null;default:false"`
	CreatedAt      time.Time `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null"`
}

func (userModel) TableName() string {
	return "users"

}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var model userModel
	err := u.db.WithContext(ctx).Where("username = ?", username).First(&model).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrUserNotFound
		}
		return nil, fmt.Errorf("query user by username failed : %w", err)
	}

	return toDomainUser(&model), nil

}

func (u *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	var model userModel
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&model).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrUserNotFound
		}
		return nil, fmt.Errorf("query user by id failed : %w", err)
	}

	return toDomainUser(&model), nil
}

func (u *UserRepository) Create(ctx context.Context, user *domain.User) error {
	model := toUserModel(user)

	if err := u.db.WithContext(ctx).Create(&model).Error; err != nil {
		if isUniqueViolation(err) {
			return repository.ErrUsernameExists
		}
		return fmt.Errorf("create user failed: %w", err)
	}

	*user = *toDomainUser(&model)
	return nil
}

func toDomainUser(model *userModel) *domain.User {
	return &domain.User{
		ID:             model.ID,
		Username:       model.Username,
		HashedPassword: model.HashedPassword,
		IsAdmin:        model.IsAdmin,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
	}
}

func toUserModel(user *domain.User) userModel {
	return userModel{
		ID:             user.ID,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		IsAdmin:        user.IsAdmin,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

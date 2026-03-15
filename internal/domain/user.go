package domain

import "time"

type User struct {
	ID             int64
	Username       string
	HashedPassword string
	IsAdmin        bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

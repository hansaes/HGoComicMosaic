package domain

import "time"

type SiteSetting struct {
	ID        int64
	Key       string
	Value     map[string]any
	CreatedAt time.Time
	UpdatedAt time.Time
}

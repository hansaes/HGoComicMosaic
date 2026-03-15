package domain

import "time"

type ResourceStatus string

const (
	ResourceStatusPending  ResourceStatus = "PENDING"
	ResourceStatusApproved ResourceStatus = "APPROVED"
	ResourceStatusRejected ResourceStatus = "REJECTED"
)

type MediaType string

const (
	MediaTypeMovie MediaType = "movie"
	MediaTypeTV    MediaType = "tv"
)

type Resource struct {
	ID           int64
	Title        string
	TitleEn      string
	Description  string
	ResourceType string
	Status       ResourceStatus

	PosterImage *string
	TmdbID      *int64
	MediaType   *MediaType

	LikesCount int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

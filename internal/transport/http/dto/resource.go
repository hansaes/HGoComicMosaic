package dto

type CreateResourceRequest struct {
	Title        string `json:"title" binding:"required"`
	TitleEn      string `json:"title_en" `
	Description  string `json:"description" binding:"required"`
	ResourceType string `json:"resource_type" binding:"required"`
}

type UpdateResourceRequest struct {
	Title        *string `json:"title" `
	TitleEn      *string `json:"title_en" `
	Description  *string `json:"description" `
	ResourceType *string `json:"resource_type" `
	PosterImage  *string `json:"poster_image"`
	TmdbId       *int64  `json:"tmdb_id"`
}

type ResourceResponse struct {
	ID           int64   `json:"id"`
	Title        string  `json:"title" `
	TitleEn      string  `json:"title_en" `
	Description  string  `json:"description" `
	ResourceType string  `json:"resource_type" `
	Status       string  `json:"status"`
	PosterImage  *string `json:"poster_image"`
	LikesCount   int     `json:"likes_count"`
}

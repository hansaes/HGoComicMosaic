package dto

type CreateUserRequest struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	IsAdmin  bool   `json:"is_admin,omitempty"`
}

type UserResponse struct {
	ID        int64  `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	IsAdmin   bool   `json:"is_admin,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

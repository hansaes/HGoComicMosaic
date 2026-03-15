package dto

type LoginRequest struct {
	Username string `json:"username,omitempty" binding:"required,min=3,max=32"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type CurrentUserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

package dto

type user struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserIdRequest struct {
	UserId string `json:"user_id"`
}

type TokenRequest struct {
	Token string `json:"token"`
}

type TokenResponse struct {
	ID        int    `json:"id"`
	UserId    string `json:"user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type AuthResponse struct {
	User         user   `json:"user"`
	TokenJwt     string `json:"token_jwt"`
	TokenRefresh string `json:"token_refresh"`
}

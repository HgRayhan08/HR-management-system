package dto

import "database/sql"

type User struct {
	Id        string       `json:"id"`
	Email     string       `json:"email"`
	Role      string       `json:"role"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   string `json:"role_id,omitempty"`
}

type UserIdRequest struct {
	UserId string `json:"user_id"`
}

type TokenRequest struct {
	Token string `json:"token"`
}

type TokenResponse struct {
	Id        string       `json:"id"`
	UserId    string       `json:"user_id"`
	Token     string       `json:"token"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at,omitempty"`
}

type AuthResponse struct {
	User         User   `json:"user"`
	TokenJwt     string `json:"token_jwt"`
	TokenRefresh string `json:"token_refresh"`
}

type RegisterResponse struct {
	ID        string       `json:"id"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type RoleRequest struct {
	ID          string `json:id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleIdRequest struct {
	RoleId string `json:"role_id"`
}

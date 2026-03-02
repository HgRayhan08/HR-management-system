package domain

import (
	"back-end-hr-manajement/dto"
	"context"
)

type UserDomain struct {
	ID        int    `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	IsActive  bool   `db:"is_active"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type RefreshTokenDomain struct {
	ID        int    `db:"id"`
	Token     string `db:"token"`
	UserId    int    `db:"user_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type UserRepository interface {
	// refresh Token
	SaveTokenRefresh(ctx context.Context, tokenData RefreshTokenDomain) error
	FindToken(ctx context.Context, userId string) (dto.TokenResponse, error)
	DeleteTokenRefresh(ctx context.Context, userId string) error
	UpdateTokenRefresh(ctx context.Context, userId string, tokenData string, time string) error
	// user
	FindtUserByEmail(ctx context.Context, email string) (UserDomain, error)
	SaveUser(ctx context.Context, user UserDomain) error
}

package domain

import (
	"back-end-hr-manajement/dto"
	"context"
	"database/sql"
)

type UserDomain struct {
	ID        string       `db:"id"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	RoleId    string       `db:"role_id"`
	IsActive  bool         `db:"is_active"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type RefreshTokenDomain struct {
	ID        string       `db:"id"`
	Token     string       `db:"token"`
	UserId    string       `db:"user_id"`
	CreatedAt sql.NullTime `db:"created_at"`
}

type RoleDomain struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type UserRepository interface {
	// refresh Token
	SaveTokenRefresh(ctx context.Context, tokenData RefreshTokenDomain) error
	FindTokenByUserId(ctx context.Context, userId string) (dto.TokenResponse, error)
	FindToken(ctx context.Context, token string) (dto.TokenResponse, error)
	DeleteTokenRefresh(ctx context.Context, userId string) error
	UpdateTokenRefresh(ctx context.Context, userId string, tokenData string, time sql.NullTime) error
	// roles
	FindRole(ctx context.Context, roleId dto.RoleIdRequest) (RoleDomain, error)
	FindAllRole(ctx context.Context) ([]RoleDomain, error)
	SaveRole(ctx context.Context, role RoleDomain) error
	DeleteRole(ctx context.Context, roleId dto.RoleIdRequest) error
	// user
	FindtUserByEmail(ctx context.Context, email string) (UserDomain, error)
	SaveUser(ctx context.Context, user UserDomain) error
}

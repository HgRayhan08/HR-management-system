package domain

import (
	"back-end-hr-manajement/dto"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
	Register(ctx context.Context, req dto.AuthRequest) error
	RefreshToken(ctx context.Context, req dto.TokenRequest) (dto.TokenResponse, error)
	Logout(ctx context.Context) error
}

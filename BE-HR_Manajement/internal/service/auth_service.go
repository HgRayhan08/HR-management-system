package service

import (
	"back-end-hr-manajement/domain"
	"back-end-hr-manajement/dto"
	"context"
)

type authService struct {
	repository domain.AuthService
}

func NewAuthService(repository domain.AuthService) domain.AuthService {
	return &authService{
		repository: repository,
	}
}

// Login implements [domain.AuthService].
func (a *authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	panic("unimplemented")
}

// Logout implements [domain.AuthService].
func (a *authService) Logout(ctx context.Context) error {
	panic("unimplemented")
}

// RefreshToken implements [domain.AuthService].
func (a *authService) RefreshToken(ctx context.Context, req dto.TokenRequest) (dto.TokenResponse, error) {
	panic("unimplemented")
}

// Register implements [domain.AuthService].
func (a *authService) Register(ctx context.Context, req dto.AuthRequest) error {
	panic("unimplemented")
}

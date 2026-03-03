package domain

import (
	"back-end-hr-manajement/dto"
	"context"

	"github.com/gofiber/fiber/v3"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
	RefreshToken(ctx context.Context, req dto.TokenRequest, f fiber.Ctx) (dto.TokenResponse, error)

	Logout(ctx context.Context, f fiber.Ctx) error
	// role
	CreateRole(ctx context.Context, req dto.RoleRequest) error
	FindRole(ctx context.Context, roleId dto.RoleIdRequest) (RoleDomain, error)
	FindAllRole(ctx context.Context) ([]RoleDomain, error)
	DeleteRole(ctx context.Context, roleId dto.RoleIdRequest) error
}

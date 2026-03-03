package service

import (
	"back-end-hr-manajement/domain"
	"back-end-hr-manajement/dto"
	"back-end-hr-manajement/internal/config"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf       *config.Config
	repository domain.UserRepository
}

func NewAuthService(repository domain.UserRepository) domain.AuthService {
	return &authService{
		repository: repository,
	}
}

// Register implements [domain.AuthService].
func (a *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	chekEmail, err := a.repository.FindtUserByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return dto.RegisterResponse{}, fiber.NewError(fiber.StatusInternalServerError, "gagal memeriksa email")
	}

	if chekEmail.ID != "" {
		return dto.RegisterResponse{}, fiber.NewError(fiber.StatusBadRequest, "email sudah terdaftar")
	}
	generatePass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, errors.New("Failed Generate Password")
	}
	user := domain.UserDomain{
		ID:        uuid.NewString(),
		Email:     req.Email,
		Password:  string(generatePass),
		RoleId:    req.RoleId,
		IsActive:  true,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
		UpdatedAt: sql.NullTime{Valid: false},
	}
	err = a.repository.SaveUser(ctx, user)
	if err != nil {
		return dto.RegisterResponse{}, errors.New("Failed Registrasi User")
	}
	return dto.RegisterResponse{
		Id:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

// Login implements [domain.AuthService].
func (a *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error) {
	chekUser, err := a.repository.FindtUserByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return dto.AuthResponse{}, fiber.NewError(fiber.StatusInternalServerError, "gagal memeriksa email")
	}

	if chekUser.ID != "" {
		return dto.AuthResponse{}, fiber.NewError(fiber.StatusBadRequest, "email sudah terdaftar")
	}

	tokenRefresh := uuid.NewString()
	err = bcrypt.CompareHashAndPassword([]byte(chekUser.Password), []byte(req.Password))

	if err != nil {
		return dto.AuthResponse{}, fiber.NewError(fiber.StatusBadRequest, "password salah")
	}
	// chek role
	checkRole, err := a.repository.FindRole(ctx, dto.RoleIdRequest{RoleId: chekUser.RoleId})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return dto.AuthResponse{}, fiber.NewError(fiber.StatusInternalServerError, "gagal memeriksa role")
	}

	// check token refresh
	dataChekToken, err := a.repository.FindTokenByUserId(ctx, chekUser.ID)
	if dataChekToken.Id != "" {
		err = a.repository.DeleteTokenRefresh(ctx, chekUser.ID)
		if err != nil {
			return dto.AuthResponse{}, fiber.NewError(fiber.StatusInternalServerError, "Failed Logout akun")
		}
	}
	// data token refresh
	tokenData := domain.RefreshTokenDomain{
		Id:        uuid.NewString(),
		UserId:    chekUser.ID,
		Token:     tokenRefresh,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	// savew token refresh
	err = a.repository.SaveTokenRefresh(ctx, tokenData)

	if err != nil {
		return dto.AuthResponse{}, fiber.NewError(fiber.StatusInternalServerError, "Failed Login User")
	}

	claim := jwt.MapClaims{
		"id":  chekUser.ID,
		"exp": time.Now().Add(time.Duration(a.conf.JWT.Exp) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.JWT.Secret))

	return dto.AuthResponse{
		User: dto.User{
			Id:        chekUser.ID,
			Email:     chekUser.Email,
			Role:      checkRole.Name,
			CreatedAt: chekUser.CreatedAt,
			UpdatedAt: chekUser.UpdatedAt,
		},
		TokenJwt:     tokenStr,
		TokenRefresh: tokenRefresh,
	}, nil
}

// Logout implements [domain.AuthService].
func (a *authService) Logout(ctx context.Context, f fiber.Ctx) error {
	userId := f.Locals("user_id")
	if userId == nil {
		return fiber.ErrUnauthorized
	}
	_, err := a.repository.FindToken(ctx, userId.(string))
	if err != nil {
		return err
	}
	return a.repository.DeleteTokenRefresh(ctx, userId.(string))
}

// RefreshToken implements [domain.AuthService].
func (a *authService) RefreshToken(ctx context.Context, req dto.TokenRequest, f fiber.Ctx) (result dto.TokenResponse, err error) {
	userData, err := a.repository.FindToken(ctx, req.Token)
	if err != nil {
		return dto.TokenResponse{}, err
	}
	claim := jwt.MapClaims{
		"id":  userData.UserId,
		"exp": time.Now().Add(time.Duration(a.conf.JWT.Exp) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(a.conf.JWT.Secret))
	if err != nil {
		return dto.TokenResponse{}, errors.New("Failed Generate Token")
	}
	err = a.repository.UpdateTokenRefresh(ctx, userData.UserId, tokenString, sql.NullTime{Time: time.Now(), Valid: true})
	if err != nil {
		return dto.TokenResponse{}, errors.New("Failed Update Token")
	}
	return dto.TokenResponse{UserId: userData.UserId, Token: tokenString, UpdatedAt: sql.NullTime{Valid: true, Time: time.Now()}}, nil
}

// CreateRole implements [domain.AuthService].
func (a *authService) CreateRole(ctx context.Context, req dto.RoleRequest) error {
	data := domain.RoleDomain{
		Id:          uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
	}
	err := a.repository.SaveRole(ctx, data)
	if err != nil {
		return errors.New("Failed Crate Role")
	}
	return nil
}

// DeleteRole implements [domain.AuthService].
func (a *authService) DeleteRole(ctx context.Context, roleId dto.RoleIdRequest) error {
	err := a.repository.DeleteRole(ctx, roleId)
	if err != nil {
		return errors.New("Failed Delete Role")
	}
	return nil
}

// FindAllRole implements [domain.AuthService].
func (a *authService) FindAllRole(ctx context.Context) ([]domain.RoleDomain, error) {
	data, err := a.repository.FindAllRole(ctx)
	if err != nil {
		return nil, errors.New("Failed Find All Role")
	}
	return data, nil
}

// FindRole implements [domain.AuthService].
func (a *authService) FindRole(ctx context.Context, roleId dto.RoleIdRequest) (domain.RoleDomain, error) {
	data, err := a.repository.FindRole(ctx, roleId)
	if err != nil {
		return domain.RoleDomain{}, errors.New("Failed Find Role")
	}
	return data, nil
}

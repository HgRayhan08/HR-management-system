package api

import (
	"back-end-hr-manajement/domain"
	"back-end-hr-manajement/dto"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
)

type authApi struct {
	authService domain.AuthService
}

func NewAuthApi(app *fiber.App, authService domain.AuthService, middlewware fiber.Handler) {
	api := authApi{authService: authService}

	app.Post("/login", api.Login)
	app.Post("/register", middlewware, api.Register)
	app.Post("/logout", middlewware, api.Logout)
	app.Post("/refresh-token", middlewware, api.RefreshToken)

	app.Get("/roles", api.GetAllRole)
	app.Get("/role", api.getRole)
	app.Post("/role", api.createRole)
	app.Delete("/role", api.deleteRole)
}

func (a *authApi) Login(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.LoginRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Request body tidak valid atau format JSON salah",
		})
	}
	res, err := a.authService.Login(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.Response{
		Code:    http.StatusOK,
		Message: "Login Berhasil",
		Data:    res,
	})
}

func (a *authApi) Register(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.RegisterRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Request body tidak valid atau format JSON salah",
		})
	}
	res, err := a.authService.Register(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.Response{
		Code:    http.StatusOK,
		Message: "Register Berhasil",
		Data:    res,
	})
}

func (a *authApi) Logout(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	err := a.authService.Logout(c, ctx)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.Response{
		Code:    http.StatusOK,
		Message: "Logout Berhasil",
	})
}

func (a *authApi) RefreshToken(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.TokenRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Request body tidak valid atau format JSON salah",
		})
	}
	res, err := a.authService.RefreshToken(c, req, ctx)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.Response{
		Code:    http.StatusOK,
		Message: "Refresh Token Berhasil",
		Data:    res,
	})
}

func (a *authApi) GetAllRole(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	result, err := a.authService.FindAllRole(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.Response{
		Code:    http.StatusOK,
		Message: "Succses Get All Role",
		Data:    result,
	})
}

func (a *authApi) createRole(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.RoleRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Request body tidak valid atau format JSON salah",
		})
	}
	err := a.authService.CreateRole(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.Response{
		Code:    http.StatusOK,
		Message: "Succses Create Role",
	})
}

func (a *authApi) getRole(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.RoleIdRequest

	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Request body tidak valid atau format JSON salah",
		})
	}
	result, err := a.authService.FindRole(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.Response{
		Code:    http.StatusOK,
		Message: "Succses Get Role",
		Data:    result,
	})

}

func (a *authApi) deleteRole(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	var req dto.RoleIdRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Request body tidak valid atau format JSON salah",
		})
	}
	err := a.authService.DeleteRole(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.Response{
		Code:    http.StatusOK,
		Message: "Succses Delete Role",
	})
}

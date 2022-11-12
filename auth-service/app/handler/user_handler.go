package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/auth-service/entity"
	"github.com/saufiroja/auth-service/interfaces"
	"github.com/saufiroja/auth-service/rpc/pb/user"
)

type UserHandler struct {
	user.UnimplementedAuthServiceServer
	UserService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	data := entity.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := h.UserService.Register(&data)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResponse{
		Status: fiber.StatusCreated,
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	token, err := h.UserService.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &user.LoginResponse{
		AccessToken: token,
	}, nil
}

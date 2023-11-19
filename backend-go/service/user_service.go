package service

import (
	"context"
	"github.com/arioprima/blog_web/models/web/request"
	"github.com/arioprima/blog_web/models/web/response"
)

type UserService interface {
	Login(ctx context.Context, request request.UserLoginRequest) (response.LoginResponse, error)
	Create(ctx context.Context, request request.UserCreateRequest) (response.UserResponse, error)
	Update(ctx context.Context, request request.UserUpdateRequest) (response.UserResponse, error)
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) (response.UserResponse, error)
	FindByUserName(ctx context.Context, username string) (response.UserResponse, error)
	FindByEmail(ctx context.Context, email string) (response.UserResponse, error)
	FindAll(context.Context) ([]response.UserResponse, error)
}

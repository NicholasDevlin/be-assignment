package service

import (
	"bank/dto/user"
	"context"
)

type UserService interface {
	Create(ctx context.Context, input user.UserRequest) (user.UserResponse, error)
	Update(ctx context.Context, input user.UserRequest) (user.UserResponse, error)
	Delete(ctx context.Context, userId string) error
	FindAll(ctx context.Context) []user.UserResponse
	FindById(ctx context.Context, userId string) (user.UserResponse, error)
}

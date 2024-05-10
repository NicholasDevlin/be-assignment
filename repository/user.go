package repository

import (
	"bank/model"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User) (model.User, error)
	Login(ctx context.Context, user model.User) (model.User, error)
	Update(ctx context.Context, user model.User) (model.User, error)
	Delete(ctx context.Context, userId string) error
	FindById(ctx context.Context, userId string) (model.User, error)
	FindAll(ctx context.Context) []model.User
}

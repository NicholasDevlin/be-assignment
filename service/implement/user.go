package serviceimplement

import (
	"bank/dto/user"
	errors "bank/helper/error"
	"bank/model"
	"bank/repository"
	"bank/service"
	"context"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

// Create implements service.UserService.
func (u *UserServiceImpl) Create(ctx context.Context, input user.UserRequest) (user.UserResponse, error) {
	userData := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	u.UserRepository.Save(ctx, userData)
	return user.UserResponse{
		Id: userData.Id,
		Email: userData.Email,
		Name: userData.Name,
	}, nil
}

// Delete implements service.UserService.
func (u *UserServiceImpl) Delete(ctx context.Context, userId string) error {
	user, err := u.UserRepository.FindById(ctx, userId)
	errors.ErrorPanic(err)
	u.UserRepository.Delete(ctx, user.Id)
	return nil
}

// FindAll implements service.UserService.
func (u *UserServiceImpl) FindAll(ctx context.Context) []user.UserResponse {
	users := u.UserRepository.FindAll(ctx)
	var userRes []user.UserResponse
	for _, v := range users {
		userRes = append(userRes, user.UserResponse{
			Id:    v.Id,
			Name:  v.Name,
			Email: v.Email,
		})
	}
	return userRes
}

// FindById implements service.UserService.
func (u *UserServiceImpl) FindById(ctx context.Context, userId string) (user.UserResponse, error) {
	userData, err := u.UserRepository.FindById(ctx, userId)
	errors.ErrorPanic(err)
	return user.UserResponse{
		Id:    userData.Id,
		Email: userData.Email,
		Name:  userData.Name,
	}, nil
}

// Update implements service.UserService.
func (u *UserServiceImpl) Update(ctx context.Context, input user.UserRequest) (user.UserResponse, error) {
	userData := model.User{
		Id:       input.Id,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	u.UserRepository.Update(ctx, userData)
	return user.UserResponse{
		Id:    userData.Id,
		Email: userData.Email,
		Name: userData.Name,
	}, nil
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

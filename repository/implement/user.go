package repositoryimplement

import (
	errors "bank/helper/error"
	"bank/model"
	"bank/prisma/db"
	"bank/repository"
	"context"
)

type UserRepositoryImpl struct {
	DB *db.PrismaClient
}

// Delete implements repository.UserRepository.
func (u *UserRepositoryImpl) Delete(ctx context.Context, userId string) error {
	_, err := u.DB.User.FindUnique(
		db.User.ID.Equals(userId),
	).Delete().Exec(ctx)
	if err != nil {
		errors.ErrorPanic(err)
		return err
	}
	return nil
}

// FindAll implements repository.UserRepository.
func (u *UserRepositoryImpl) FindAll(ctx context.Context) []model.User {
	allUser, err := u.DB.User.FindMany().Exec(ctx)
	errors.ErrorPanic(err)

	var users []model.User
	for _, user := range allUser {
		users = append(users, model.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return users
}

// FindById implements repository.UserRepository.
func (u *UserRepositoryImpl) FindById(ctx context.Context, userId string) (model.User, error) {
	user, err := u.DB.User.FindFirst(db.User.ID.Equals(userId)).Exec(ctx)
	if err != nil {
		errors.ErrorPanic(err)
		return model.User{}, err
	}

	if user == nil {
		return model.User{}, errors.ERR_USER_NOT_FOUND
	}

	return model.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// Save implements repository.UserRepository.
func (u *UserRepositoryImpl) Save(ctx context.Context, user model.User) (model.User, error) {
	result, err := u.DB.User.CreateOne(
		db.User.Name.Set(user.Name),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
	).Exec(ctx)
	if err != nil {
		errors.ErrorPanic(err)
		return model.User{}, err
	}
	return model.User{
		Id:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}, nil
}

// Update implements repository.UserRepository.
func (u *UserRepositoryImpl) Update(ctx context.Context, user model.User) (model.User, error) {
	result, err := u.DB.User.FindUnique(db.User.ID.Equals(user.Id)).Update(
		db.User.Name.Set(user.Name),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
	).Exec(ctx)
	if err != nil {
		errors.ErrorPanic(err)
		return model.User{}, err
	}

	return model.User{
		Id:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}, nil
}

func NewUserRepository(DB *db.PrismaClient) repository.UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

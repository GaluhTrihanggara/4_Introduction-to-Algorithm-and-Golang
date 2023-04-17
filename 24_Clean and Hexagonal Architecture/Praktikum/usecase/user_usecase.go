package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"

	"github.com/pkg/errors"
)

type UserUsecase interface {
	CreateUser(user *model.User) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) CreateUser(user *model.User) (*model.User, error) {
	createdUser, err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	return createdUser, nil
}

func (u *userUsecase) GetAllUsers() ([]*model.User, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all users")
	}
	return users, nil
}

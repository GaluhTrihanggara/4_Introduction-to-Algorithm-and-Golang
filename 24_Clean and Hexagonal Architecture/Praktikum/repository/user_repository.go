package repository

import (
	"belajar-go-echo/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *model.User) (*model.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all users")
	}
	return users, nil
}

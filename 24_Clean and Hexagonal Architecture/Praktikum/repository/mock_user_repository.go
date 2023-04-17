package repository

import (
	"belajar-go-echo/model"

	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

// CreateUser implements UserRepository
func (m *MockUserRepository) CreateUser(user *model.User) (*model.User, error) {
	args := m.Called(user)
	return args.Get(0).(*model.User), args.Error(1)
}

// GetAllUsers implements UserRepository
func (*MockUserRepository) GetAllUsers() ([]*model.User, error) {
	panic("unimplemented")
}

// NewMockUserRepository creates a new instance of MockUserRepository
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

// Find is a mock implementation of the Find method in UserRepository
func (m *MockUserRepository) Find(email string) (model.User, error) {
	ret := m.Called(email)
	return ret.Get(0).(model.User), ret.Error(1)
}

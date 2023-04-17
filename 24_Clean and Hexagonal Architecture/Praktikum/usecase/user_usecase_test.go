package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserUsecase(t *testing.T) {
	mockUserRepo := &repository.MockUserRepository{} // Ganti dengan mock repository sesuai kebutuhan
	type args struct {
		userRepo repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want UserUsecase
	}{
		{
			name: "Test NewUserUsecase",
			args: args{
				userRepo: mockUserRepo,
			},
			want: NewUserUsecase(mockUserRepo),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserUsecase(tt.args.userRepo)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser_Success(t *testing.T) {
	// Prepare mock UserRepository
	mockUserRepo := new(repository.MockUserRepository)
	defer mockUserRepo.AssertExpectations(t)

	// Prepare userUsecase with mock UserRepository
	userUsecase := NewUserUsecase(mockUserRepo)

	// Prepare input and expected output
	inputUser := &model.User{
		Name:     "John",
		Email:    "john@example.com",
		Password: "password",
	}
	expectedUser := &model.User{
		ID:       1,
		Name:     "John",
		Email:    "john@example.com",
		Password: "password",
	}

	// Mock UserRepository's CreateUser to return expected output
	mockUserRepo.On("CreateUser", inputUser).Return(expectedUser, nil)

	// Call CreateUser on userUsecase
	createdUser, err := userUsecase.CreateUser(inputUser)

	// Assert that CreateUser returned expected output
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, createdUser)
}

func TestCreateUser_Error(t *testing.T) {
	// Prepare mock UserRepository
	mockUserRepo := new(repository.MockUserRepository)
	defer mockUserRepo.AssertExpectations(t)

	// Prepare userUsecase with mock UserRepository
	userUsecase := NewUserUsecase(mockUserRepo)

	// Prepare input
	inputUser := &model.User{
		Name:     "John",
		Email:    "john@example.com",
		Password: "password",
	}

	// Mock UserRepository's CreateUser to return error
	expectedError := errors.New("failed to create user")
	mockUserRepo.On("CreateUser", inputUser).Return(nil, expectedError)

	// Call CreateUser on userUsecase
	createdUser, err := userUsecase.CreateUser(inputUser)

	// Assert that CreateUser returned expected error
	assert.Error(t, err)
	assert.Nil(t, createdUser)
	assert.EqualError(t, err, "failed to create user")
}

func Test_userUsecase_GetAllUsers(t *testing.T) {
	mockUserRepo := &repository.MockUserRepository{} // Ganti dengan mock repository sesuai kebutuhan
	tests := []struct {
		name    string
		u       *userUsecase
		want    []*model.User
		wantErr bool
	}{
		{
			name: "Test GetAllUsers Success",
			u: &userUsecase{
				userRepo: mockUserRepo,
			},
			want: []*model.User{
				{
					ID:   1,
					Name: "John Doe",
				},
				{
					ID:   2,
					Name: "Jane Smith",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetAllUsers()
			if (err !=
				nil) != tt.wantErr {
				t.Errorf("userUsecase.GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

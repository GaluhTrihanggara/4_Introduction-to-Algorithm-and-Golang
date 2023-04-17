package controller

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	// Bind request data to DTO
	req := new(dto.UserCreateRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate request data
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// Create user model from request data
	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	// Call use case to create user
	createdUser, err := c.userUsecase.CreateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// Create response DTO
	res := &dto.UserResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *UserController) GetAllUsers(ctx echo.Context) error {
	// Call use case to get all users
	users, err := c.userUsecase.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// Create response DTO
	var res []*dto.UserResponse
	for _, user := range users {
		res = append(res, &dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
		})
	}

	return ctx.JSON(http.StatusOK, res)
}

package main

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// Create repository
	userRepository := repository.NewUserRepository(db)

	// Create usecase
	userUsecase := usecase.NewUserUsecase(userRepository)

	// Create controller
	userController := controller.NewUserController(userUsecase)

	// Define routes
	e.POST("/users", userController.CreateUser)
	e.GET("/users", userController.GetAllUsers)
}

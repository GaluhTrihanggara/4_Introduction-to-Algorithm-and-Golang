package routes

import (
	"Praktikum/constants"
	"Praktikum/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	jwtMiddleware := middleware.JWT([]byte(constants.SECRET_JWT))
	// Route / to handler function
	e.GET("/users", controllers.GetUsersController, jwtMiddleware)
	e.GET("/users/:id", controllers.GetUserController, jwtMiddleware)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController, jwtMiddleware)
	e.PUT("/users/:id", controllers.UpdateUserController, jwtMiddleware)
	// start the server, and log if it fails
	return e
}

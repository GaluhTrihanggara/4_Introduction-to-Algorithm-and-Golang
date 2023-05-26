package routes

import (
	"codecompetence/constants"
	"codecompetence/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	jwtMiddleware := middleware.JWT([]byte(constants.SECRET_JWT))
	// Route /users to handler functions for user authentication and authorization
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	// Route /users/:id to handler functions for updating and deleting a user
	userGroup := e.Group("/users", jwtMiddleware)
	userGroup.GET("/:id", controllers.GetUserController)
	userGroup.PUT("/:id", controllers.UpdateUserController)
	userGroup.DELETE("/:id", controllers.DeleteUserController)
	userGroup.GET("/items/:id", controllers.GetItemController)

	// Route /users to handler functions for user authentication and authorization
	e.POST("/admin/login", controllers.LoginAdminController)
	e.POST("/admin", controllers.CreateAdminController)

	// group routes for admin
	adminGroup := e.Group("/admin", jwtMiddleware)
	adminGroup.GET("", controllers.GetAdminsController)
	adminGroup.GET("/:id", controllers.GetAdminController)
	adminGroup.PUT("/:id", controllers.UpdateAdminController)
	adminGroup.DELETE("/:id", controllers.DeleteAdminController)

	// Route /users to handler function
	adminGroup.GET("/users", controllers.GetUsersController)
	adminGroup.GET("/users/:id", controllers.GetUserController)
	adminGroup.PUT("/users/:id", controllers.UpdateUserController)
	adminGroup.DELETE("/users/:id", controllers.DeleteUserController)

	// Route /items to handler function
	adminGroup.POST("/items", controllers.CreateItemController)
	adminGroup.GET("/items", controllers.GetItemsController)
	adminGroup.GET("/items/:id", controllers.GetItemController)
	adminGroup.GET("/items/category/:category_id", controllers.GetItemsByCategoryIDController)
	adminGroup.GET("/item?keyword=item_name", controllers.GetItemsByKeywordController)
	adminGroup.PUT("/items/:id", controllers.UpdateItemController)
	adminGroup.DELETE("/items/:id", controllers.DeleteItemController)

	// Route /category to handler function
	adminGroup.POST("/category", controllers.CreateCategoryController)
	adminGroup.GET("/categories", controllers.GetCategoriesController)
	adminGroup.GET("/category/:id", controllers.GetCategoryController)
	adminGroup.PUT("/category/:id", controllers.UpdateCategoryController)
	adminGroup.DELETE("/category/:id", controllers.DeleteCategoryController)

	return e
}

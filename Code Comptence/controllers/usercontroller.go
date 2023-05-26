package controllers

import (
	"codecompetence/config"
	"codecompetence/lib/database"
	"codecompetence/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

// get all users
func GetUsersController(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "method not allowed")
	}

	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success get user by id",
		"data_user": user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "user not available",
		})
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal hapus data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil menghapus data",
		"user":   user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	body := new(models.User)

	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id parameter",
		})
	}

	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "user not available",
		})
	}

	user.Id = body.Id
	user.Nama = body.Nama
	user.Username = body.Username
	user.Email = body.Email
	user.Password = body.Password
	user.Alamat = body.Alamat
	user.Handphone = body.Handphone

	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil update data",
		"user":   user,
	})
}

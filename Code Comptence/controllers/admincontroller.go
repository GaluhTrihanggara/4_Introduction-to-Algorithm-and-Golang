package controllers

import (
	"codecompetence/config"
	"codecompetence/lib/database"
	"codecompetence/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	admins, e := database.LoginAdmin(&admin)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"admins": admins,
	})
}

// get all admins
func GetAdminsController(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "method not allowed")
	}

	var admins []models.Admin

	if err := config.DB.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all admins",
		"admins":  admins,
	})
}

// get admin by id
func GetAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var admin models.Admin
	if err := config.DB.First(&admin, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success get admin by id",
		"data_user": admin,
	})
}

// create new admin
func CreateAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	if err := config.DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new admin",
		"admin":   admin,
	})
}

// delete admin by id
func DeleteAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	var admin models.Admin
	if err := config.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "admin not available",
		})
	}

	if err := config.DB.Delete(&admin).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal hapus data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil menghapus data",
		"admin":  admin,
	})
}

// update admin by id
func UpdateAdminController(c echo.Context) error {
	body := new(models.Admin)

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

	var admin models.Admin
	if err := config.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "admin tidak tersedia",
		})
	}
	admin.Username = body.Username
	admin.Password = body.Password

	if err := config.DB.Save(&admin).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal memperbarui data admin",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil memperbarui data admin",
		"admin":  admin,
	})
}

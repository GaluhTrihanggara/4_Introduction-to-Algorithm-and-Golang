package controllers

import (
	"codecompetence/lib/database"
	"codecompetence/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateCategoryController creates a new category
func CreateCategoryController(c echo.Context) error {
	category := models.Category{}
	c.Bind(&category)
	if err := database.CreateCategory(&category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new category",
		"category": category,
	})
}

// GetCategoriesController retrieves all categories
func GetCategoriesController(c echo.Context) error {
	categories, err := database.GetCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success get all categories",
		"categories": categories,
	})
}

// GetCategoryController retrieves a category by its ID
func GetCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	category, err := database.GetCategoryByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get category by id",
		"category": category,
	})
}

// UpdateCategoryController updates a category by its ID
func UpdateCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	category, err := database.GetCategoryByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	body := new(models.Category)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid request body",
		})
	}

	category.Name = body.Name

	if err := database.UpdateCategory(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to update category",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success update category",
		"category": category,
	})
}

// DeleteCategoryController deletes a category by its ID
func DeleteCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	category, err := database.GetCategoryByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.DeleteCategory(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to delete category",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success delete category",
		"category": category,
	})
}

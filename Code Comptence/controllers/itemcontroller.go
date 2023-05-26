package controllers

import (
	"codecompetence/config"
	"codecompetence/lib/database"
	"codecompetence/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateItemController creates a new item
func CreateItemController(c echo.Context) error {
	item := models.Item{}
	c.Bind(&item)
	if err := config.DB.Save(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"item":    item,
	})
}

// GetItemsController retrieves all items
func GetItemsController(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "method not allowed")
	}
	var items []models.Item

	if err := config.DB.Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all items",
		"items":   items,
	})
}

// GetItemController retrieves an item by its ID
func GetItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item by id",
		"item":    item,
	})
}

func GetItemsByCategoryIDController(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}

	items, err := database.GetItemsByCategoryID(categoryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve items")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"items":   items,
	})
}

// GetItemsByKeywordController retrieves items based on the provided keyword
func GetItemsByKeywordController(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	items, err := database.GetItemsByKeyword(keyword)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get items by keyword",
		"items":   items,
	})
}

// UpdateItemController updates an item by its ID
func UpdateItemController(c echo.Context) error {
	body := new(models.Item)
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

	var item models.Item
	if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "item not available",
		})
	}

	item.Name = body.Name
	item.Category_id = body.Category_id
	item.Deskripsi = body.Deskripsi
	item.Jumlah_Stok = body.Jumlah_Stok
	item.Harga_Barang = body.Harga_Barang

	if err := config.DB.Save(&item).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal update data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "berhasil update data",
		"item":   item,
	})
}

// DeleteItemController deletes an item by its ID
func DeleteItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}
	var item models.Item
	if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "item not available",
		})
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to delete item",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete item",
		"item":   item,
	})
}

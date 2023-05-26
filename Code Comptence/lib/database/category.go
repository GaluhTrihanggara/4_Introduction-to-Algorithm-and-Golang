package database

import (
	"codecompetence/config"
	"codecompetence/models"
)

// GetCategoryByID returns a category by its ID
func GetCategoryByID(id int) (*models.Category, error) {
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateCategory(category *models.Category) error {
	if err := config.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCategory(category *models.Category) error {
	if err := config.DB.Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCategory(category *models.Category) error {
	if err := config.DB.Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

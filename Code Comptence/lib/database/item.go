package database

import (
	"codecompetence/config"
	"codecompetence/models"
)

// GetItemByID returns an item by its ID
func GetItemByID(id int) (*models.Item, error) {
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func GetItems() ([]models.Item, error) {
	var items []models.Item
	if err := config.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func GetItemsByCategoryID(categoryID int) ([]models.Item, error) {
	var items []models.Item
	if err := config.DB.Where("category_id = ?", categoryID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func GetItemsByKeyword(keyword string) ([]models.Item, error) {
	var items []models.Item
	if err := config.DB.Where("name LIKE ?", "%"+keyword+"%").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func CreateItem(item *models.Item) error {
	if err := config.DB.Create(&item).Error; err != nil {
		return err
	}
	return nil
}

func UpdateItem(item *models.Item) error {
	if err := config.DB.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItem(item *models.Item) error {
	if err := config.DB.Delete(&item).Error; err != nil {
		return err
	}
	return nil
}

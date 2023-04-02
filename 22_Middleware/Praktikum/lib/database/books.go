package database

import (
	"Praktikum/config"
	"Praktikum/models"
)

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

package database

import (
	"Praktikum/config"
	"Praktikum/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

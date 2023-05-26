package database

import (
	"codecompetence/config"
	"codecompetence/middlewares"
	"codecompetence/models"
)

func LoginAdmin(admin *models.Admin) (interface{}, error) {
	var err error
	if err = config.DB.Where("username = ? AND password = ?", admin.Username, admin.Password).First(admin).Error; err != nil {
		return nil, err
	}

	var token models.Token
	token.Token, err = middlewares.CreateToken(int(admin.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(admin).Error; err != nil {
		return nil, err
	}

	return token, nil
}

// GetAdminByID returns an admin by its ID
func GetAdminByID(id int) (*models.Admin, error) {
	var admin models.Admin
	if err := config.DB.First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func CreateAdmin(admin *models.Admin) error {
	if err := config.DB.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAdmin(admin *models.Admin) error {
	if err := config.DB.Save(&admin).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAdmin(admin *models.Admin) error {
	if err := config.DB.Delete(&admin).Error; err != nil {
		return err
	}
	return nil
}

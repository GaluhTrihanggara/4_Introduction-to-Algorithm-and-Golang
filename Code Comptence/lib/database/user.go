package database

import (
	"codecompetence/config"
	"codecompetence/middlewares"
	"codecompetence/models"
)

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	var token models.Token
	token.Token, err = middlewares.CreateToken(int(user.Id))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return token, nil
}

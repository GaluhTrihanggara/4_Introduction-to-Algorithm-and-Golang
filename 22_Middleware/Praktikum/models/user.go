package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func GetUsers() ([]User, error) {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id uint) (*User, error) {
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) (*User, error) {
	if err := DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *User) (*User, error) {
	if err := DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id uint) error {
	if err := DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

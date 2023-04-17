package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token" form:"token"`
}

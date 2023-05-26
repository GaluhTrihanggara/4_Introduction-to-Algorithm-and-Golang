package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}

package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id   int    `json:"Id"`
	Name string `json:"name" form:"name"`
}

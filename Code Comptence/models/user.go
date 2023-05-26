package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int    `json:"Id"`
	Nama      string `json:"nama" form:"nama"`
	Username  string `json:"username" form:"username"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Alamat    string `json:"alamat" form:"alamat"`
	Handphone string `json:"handphone" form:"handphone"`
}

type Token struct {
	Token string `json:"token" form:"token"`
}

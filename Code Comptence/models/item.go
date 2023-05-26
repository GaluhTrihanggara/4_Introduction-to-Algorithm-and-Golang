package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Id           int    `json:"Id"`
	Category_id  int    `json:"category_id" form:"category_id"`
	Name         string `json:"name" form:"name"`
	Deskripsi    string `json:"deskripsi" form:"deskripsi"`
	Jumlah_Stok  int    `json:"jumlah_stok" form:"jumlah_stok"`
	Harga_Barang int    `json:"harga_barang" form:"harga_barang"`
}

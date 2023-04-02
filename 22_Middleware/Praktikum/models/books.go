package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Id       uint   `json:"id" form:"id"`
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}

func GetBooks() ([]Book, error) {
	var books []Book
	if err := DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(id uint) (*Book, error) {
	var book Book
	if err := DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func CreateBook(book *Book) (*Book, error) {
	if err := DB.Create(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(book *Book) (*Book, error) {
	if err := DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id uint) error {
	if err := DB.Delete(&Book{}, id).Error; err != nil {
		return err
	}
	return nil
}

package models

import (
	"github.com/bakhodur-nazriev/my-blog/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Author struct {
	gorm.Model
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Photo    string `json:"photo"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Author{})
}

func GetAllAuthors() []Author {
	var Authors []Author
	db.Find(&Authors)
	return Authors
}

func GetAuthorById(Id int64) (*Author, *gorm.DB) {
	var getAuthor Author
	db := db.Where("ID=?", Id).Find(&getAuthor)
	return &getAuthor, db
}

func (a *Author) CreateAuthor() *Author {
	db.NewRecord(a)
	db.Create(&a)
	return a
}

func DeleteAuthor(Id int64) Author {
	var author Author
	db.Where("ID=?", Id).Delete(author)
	return author
}

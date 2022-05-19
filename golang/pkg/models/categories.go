package models

import (
	"github.com/bakhodur-nazriev/my-blog/pkg/config"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Category{})
}

func GetAllCategories() []Category {
	var Categories []Category
	db.Find(&Categories)
	return Categories
}

func GetCategoryById(Id int64) (*Category, *gorm.DB) {
	var getCategory Category
	db := db.Where("ID=?", Id).Find(&getCategory)
	return &getCategory, db
}

func (c *Category) CreateCategory() *Category {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func DeleteCategory(Id int64) Category {
	var category Category
	db.Where("ID=?").Delete(category)
	return category

}

package models

import (
	"github.com/bakhodur-nazriev/my-blog/pkg/config"
	"github.com/jinzhu/gorm"
)

type Articles struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Articles{})
}

func GetAllArticles() []Articles {
	var Articles []Articles
	db.Find(&Articles)
	return Articles
}

func GetArticleById(Id int64) (*Articles, *gorm.DB) {
	var getArticle Articles
	db := db.Where("ID=?", Id).Find(&getArticle)
	return &getArticle, db
}

func (a *Articles) CreateArticle() *Articles {
	db.NewRecord(a)
	db.Create(&a)
	return a
}

func DeleteArticle(Id int64) Articles {
	var article Articles
	db.Where("ID=?", Id).Delete(article)
	return article
}

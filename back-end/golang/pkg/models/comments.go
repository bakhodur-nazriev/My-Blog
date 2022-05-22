package models

import (
	"github.com/bakhodur-nazriev/my-blog/pkg/config"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Comment{})
}

func GetAllComments() []Comment {
	var Comments []Comment
	db.Find(&Comments)
	return Comments
}

func GetCommentById(Id int64) (*Comment, *gorm.DB) {
	var getComment Comment
	db := db.Where("ID=?", Id).Find(&getComment)
	return &getComment, db
}

func (c *Comment) CreateComment() *Comment {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func DeleteComment(Id int64) Comment {
	var comment Comment
	db.Where("ID=?", Id).Delete(comment)
	return comment
}

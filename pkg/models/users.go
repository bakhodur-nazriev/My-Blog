package models

import (
	"github.com/bakhodur-nazriev/my-blog/pkg/config"
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Users{})
}

func GetAllUsers() []Users {
	var Users []Users
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*Users, *gorm.DB) {
	var getUser Users
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func (u *Users) CreateUser() *Users {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func DeleteUser(Id int64) Users {
	var user Users
	db.Where("ID=?", Id).Delete(user)
	return user
}

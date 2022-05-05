package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/bakhodur-nazriev/my-blog/pkg/models"
	"github.com/bakhodur-nazriev/my-blog/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewArticle models.Articles

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	newArticle := models.GetAllArticles()
	res, _ := json.Marshal(newArticle)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	ID, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	articleDetails, _ := models.GetArticleById(ID)
	res, _ := json.Marshal(articleDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	CreateArticle := &models.Articles{}
	utils.ParseBody(r, CreateArticle)
	a := CreateArticle.CreateArticle()
	res, _ := json.Marshal(a)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var updateArticle = &models.Articles{}
	utils.ParseBody(r, updateArticle)
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	ID, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	articleDetails, db := models.GetArticleById(ID)
	if updateArticle.Title != "" {
		articleDetails.Title = updateArticle.Title
	}

	if updateArticle.Body != "" {
		articleDetails.Body = updateArticle.Body
	}

	db.Save(&articleDetails)
	res, _ := json.Marshal(articleDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	ID, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		fmt.Println("error while parsong")
	}
	article := models.DeleteArticle(ID)
	res, _ := json.Marshal(article)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

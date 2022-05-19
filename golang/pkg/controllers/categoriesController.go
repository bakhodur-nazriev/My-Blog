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

var NewCategory models.Category

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	newCategory := models.GetAllCategory()
	res, _ := json.Marshal(newCategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"]
	ID, err := strconv.ParseInt(categoryId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	categoryDetails, _ := models.GetCategoryById(ID)
	res, _ := json.Marshal(categoryDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	CreateCategory := &models.Category{}
	utils.ParseBody(r, CreateCategory)
	c := CreateCategory.CreateCategory()
	res, _ := json.Marshal(c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {}

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

var NewAuthor models.Author

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	newAuthor := models.GetAllAuthors()
	res, _ := json.Marshal(newAuthor)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	authorDetails, _ := models.GetAuthorById(ID)
	res, _ := json.Marshal(authorDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	CreateAuthor := &models.Author{}
	utils.ParseBody(r, CreateAuthor)
	a := CreateAuthor.CreateAuthor()
	res, _ := json.Marshal(a)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var updateAuthor = &models.Author{}
	utils.ParseBody(r, updateAuthor)
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	authorDetails, db := models.GetAuthorById(ID)
	if updateAuthor.Name != "" {
		authorDetails.Name = updateAuthor.Name
	}
	if updateAuthor.Birthday != "" {
		authorDetails.Birthday = updateAuthor.Birthday
	}
	if updateAuthor.Photo != "" {
		authorDetails.Photo = updateAuthor.Photo
	}

	db.Save(&authorDetails)
	res, _ := json.Marshal(authorDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	ID, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing!")
	}
	author := models.DeleteAuthor(ID)
	res, _ := json.Marshal(author)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

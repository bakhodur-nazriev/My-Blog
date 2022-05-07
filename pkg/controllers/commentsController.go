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

var NewComment models.Comment

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	newComment := models.GetAllComments()
	res, _ := json.Marshal(newComment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCommentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId := vars["authorId"]
	ID, err := strconv.ParseInt(commentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	commentDetails, _ := models.GetCommentById(ID)
	res, _ := json.Marshal(commentDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	CreateComment := &models.Comment{}
	utils.ParseBody(r, CreateComment)
	c := CreateComment.CreateComment()
	res, _ := json.Marshal(c)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {

}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId := vars["commentId"]
	ID, err := strconv.ParseInt(commentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	comment := models.DeleteComment(ID)
	res, _ := json.Marshal(comment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

package main

import (
	"github.com/bakhodur-nazriev/my-blog/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterMyBlogRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9020", r))
}

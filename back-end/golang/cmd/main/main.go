package main

import (
	"fmt"
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
	fmt.Println("Sever running at port:9020")
	log.Fatal(http.ListenAndServe(":9020", r))
}

package routes

import (
	"github.com/bakhodur-nazriev/my-blog/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMyBlogRoutes = func(router *mux.Router) {
	/* Authors */
	router.HandleFunc("/authors/", controllers.GetAllAuthors).Methods("GET")
	router.HandleFunc("/authors/{id}", controllers.GetAuthorById).Methods("GET")
	router.HandleFunc("/authors/", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors/{id}", controllers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{id}", controllers.DeleteAuthor).Methods("DELETE")

	/* Users */
	router.HandleFunc("/users/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	/* Articles */
	router.HandleFunc("/articles/", controllers.GetAllArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", controllers.GetArticleById).Methods("GET")
	router.HandleFunc("/articles/", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", controllers.UpdateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", controllers.DeleteArticle).Methods("DELETE")

	/* Comments */
	router.HandleFunc("/comments/", controllers.GetAllComments).Methods("GET")
	router.HandleFunc("/comments/{id}", controllers.GetCommentById).Methods("GET")
	router.HandleFunc("/comments/", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/comments/{id}", controllers.UpdateComment).Methods("PUT")
	router.HandleFunc("/comments/{id}", controllers.DeleteComment).Methods("DELETE")

	/* Categories */
	router.HandleFunc("/categories/", controllers.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", controllers.GetCategoryById).Methods("GET")
	router.HandleFunc("/categories/", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", controllers.DeleteCategory).Methods("DELETE")
}

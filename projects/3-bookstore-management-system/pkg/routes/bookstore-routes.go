package routes

import (
	"github.com/gorilla/mux"
	"github.com/saurabhsri108/learn-golang/projects/3-bookstore-management-system/pkg/controllers"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFuc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFuc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFuc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFuc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFuc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}

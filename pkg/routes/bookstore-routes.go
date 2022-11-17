package routes

import (
	"github.com/bedLad/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.C_createBook).Methods("POST")
	router.HandleFunc("/book/", controllers.C_getBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.C_getBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.C_updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.C_deleteBook).Methods("DELETE")
}

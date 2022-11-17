/*
The purpose of this file is to deliver the methods that the bookstore-routes.go calls.
This file interacts with rest of the program as its the controller.
*/
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bedLad/go-bookstore/pkg/models"
	"github.com/bedLad/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func C_getBook(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	response, _ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
	return
}

func C_getBookById(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 0)

	if err != nil {
		fmt.Println("error while conversion")
	}

	bookDetails, _ := models.GetBookById(ID)

	response, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func C_createBook(res http.ResponseWriter, req *http.Request) {
	create_book := &models.Book{}
	utils.ParseBody(req, create_book)
	b := create_book.CreateBook()
	response, _ := json.Marshal(b)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func C_deleteBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	models.DeleteBook(ID)
	response, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func C_updateBook(res http.ResponseWriter, req *http.Request) {
	var update_book = &models.Book{}  // This is the book details which will be input by the user at front end form
	utils.ParseBody(req, update_book) // The book details input by user is sent as request to server then parsed and stored in update_book
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book, db := models.GetBookById(ID)

	if update_book.Name != "" {
		book.Name = update_book.Name
	}
	if update_book.Author != "" {
		book.Author = update_book.Author
	}
	if update_book.Publication != "" {
		book.Publication = update_book.Publication
	}

	db.Save(&book)

	response, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "applicaiton/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

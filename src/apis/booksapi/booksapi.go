package booksapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"../../config"
	"../../entities"
	"../../models"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello from Books")
}

func Hi(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hi from Books")
}

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		bookModel := models.BookModel{
			DB: db,
		}
		books, err2 := bookModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, books)
		}
	}
}

// Add new book
func createBook(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		bookModel := models.BookModel{
			DB: db,
		}
		var book entities.Book
		book.ID = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&book)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := bookModel.Create(&book)
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				respondWithJson(response, http.StatusOK, book)
			}
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

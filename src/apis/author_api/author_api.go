package booksapi

import (
	"../../models"
	"encoding/json"
	"../../entities"
	"github.com/gorilla/mux"
	"net/http"
)

func FindAll(res http.ResponseWriter, request *http.Request) {
	var bookModel models.BookModel
	books, err := bookModel.FindAll()
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, books)
	}
}

func Search(res http.ResponseWriter, request *http.Request) {
	var bookModel models.BookModel
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	books, err := bookModel.Search(keyword)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {

		respondWithJson(res, http.StatusOK, books)
	}
}

func Find(res http.ResponseWriter, request *http.Request) {
	var bookModel models.BookModel
	vars := mux.Vars(request)
	id := vars["id"]
	book, err := bookModel.Find(id)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {

		respondWithJson(res, http.StatusOK, book)
	}
}

func Create(res http.ResponseWriter, request *http.Request) {
	var book entities.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var bookModel models.BookModel
		err2 := bookModel.Create(&book)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(res, http.StatusOK, book)
		}
	}
}

func Update(res http.ResponseWriter, request *http.Request) {
	var book entities.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		vars := mux.Vars(request)
		id := vars["id"]
		var bookModel models.BookModel
		err2 := bookModel.Update(id, &book)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(res, http.StatusOK, book)
		}
	}
}

func Save(res http.ResponseWriter, request *http.Request) {
	var book entities.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var bookModel models.BookModel
		err2 := bookModel.Save(&book)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(res, http.StatusOK, book)
		}
	}
}

func Delete(res http.ResponseWriter, request *http.Request) {
	var bookModel models.BookModel
	vars := mux.Vars(request)
	id := vars["id"]
	book, err := bookModel.Find(id)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		err2 := bookModel.Delete(book)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(res, http.StatusOK, book)
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

package booksapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"

	"../../config"
	"../../dao"
	"../../entities"
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
		bookDAO := dao.BookDAO{
			DB: db,
		}
		books, err2 := bookDAO.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, books)
		}
	}
}

func Find(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		bookDAO := dao.BookDAO{
			DB: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		book, err2 := bookDAO.Find(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, book)
		}
	}
}

// Add new book
func Create(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		bookDAO := dao.BookDAO{
			DB: db,
		}
		var book entities.Book
		book.Id = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&book)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := bookDAO.Create(&book)
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				respondWithJson(response, http.StatusOK, book)
			}
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		bookDAO := dao.BookDAO{
			DB: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		err2 := bookDAO.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, entities.Book{})
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		bookDAO := dao.BookDAO{
			DB: db,
		}
		var book entities.Book
		err2 := json.NewDecoder(request.Body).Decode(&book)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := bookDAO.Update(&book)
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

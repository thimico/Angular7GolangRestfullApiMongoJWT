package main

import (
	"log"
	"net/http"

	"./src/apis/accountapi"
	"./src/apis/booksapi"
	"./src/apis/jwtauth"
	"./src/middlewares/basicauth"
	"./src/middlewares/jwtauthmiddleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/generatekey", jwtauth.GenerateToken).Methods("POST")
	// r.HandleFunc("/api/account/create", accountapi.Create).Methods("POST")
	r.Handle("/api/account/create", basicauth.BasicAuth(http.HandlerFunc(accountapi.Create))).Methods("POST")
	r.Handle("/api/book/findall", jwtauthmiddleware.JWTAuth(http.HandlerFunc(booksapi.FindAll))).Methods("GET")
	r.Handle("/api/book/hi", basicauth.BasicAuth(http.HandlerFunc(booksapi.Hi))).Methods("GET")

	// Route handles & endpoints
	// r.HandleFunc("/api/books", getBooks).Methods("GET")
	// r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("/api/books", createBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

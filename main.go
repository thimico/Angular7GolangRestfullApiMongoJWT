package main

import (
	"log"
	"net/http"

	a "./src/apis/abstractapi"
	"./src/apis/accountapi"
	"./src/apis/booksapi"
	"./src/apis/jwtauth"
	. "./src/apis/productapi"
	"./src/middlewares/jwtauthmiddleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/api/login", jwtauth.GenerateToken).Methods("POST")
	r.HandleFunc("/api/user", accountapi.Create).Methods("POST")
	r.HandleFunc("/api/users", booksapi.FindAll).Methods("GET")
	r.HandleFunc("/api/user/{id}", booksapi.Find).Methods("GET")
	r.HandleFunc("/api/user/{id}", booksapi.Update).Methods("PUT")
	r.HandleFunc("/api/user/{id}", booksapi.Delete).Methods("DELETE")

	//books
	r.HandleFunc("/api/books", booksapi.FindAll).Methods("GET")
	r.HandleFunc("/api/books/{id}", booksapi.Find).Methods("GET")
	r.HandleFunc("/api/books", booksapi.Create).Methods("POST")
	r.HandleFunc("/api/books/{id}", booksapi.Update).Methods("PUT")
	r.HandleFunc("/api/books/{id}", booksapi.Delete).Methods("DELETE")

	//products
	productApi := ProductAPI{AbstractAPI: a.AbstractAPI{FirstName: "John"}}
	r.Handle("/api/products/all", jwtauthmiddleware.JWTAuth(http.HandlerFunc(productApi.ObterTodos))).Methods("GET")
	r.HandleFunc("/api/products", productApi.ObterTodos).Methods("GET")
	r.HandleFunc("/api/products", productApi.Create).Methods("POST")
	r.HandleFunc("/api/products/{id}", productApi.Find).Methods("GET")
	r.HandleFunc("/api/products/{id}", productApi.Delete).Methods("DELETE")
	r.HandleFunc("/api/products/{id}", productApi.Update).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}

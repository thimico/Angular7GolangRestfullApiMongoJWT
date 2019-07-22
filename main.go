package main

import (
	"./src/apis/product_api"
	"./src/apis/booksapi"
	"./src/apis/author_api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//books
	r.HandleFunc("/api/books", booksapi.FindAll).Methods("GET")
	r.HandleFunc("/api/books/{id}", booksapi.Find).Methods("GET")
	r.HandleFunc("/api/books/search/{keyword}", booksapi.Search).Methods("GET")
	r.HandleFunc("/api/books", booksapi.Create).Methods("POST")
	r.HandleFunc("/api/books/{id}", booksapi.Update).Methods("PUT")
	r.HandleFunc("/api/books", booksapi.Save).Methods("PUT")
	r.HandleFunc("/api/books/{id}", booksapi.Delete).Methods("DELETE")

	//author
	r.HandleFunc("/api/authors", author_api.FindAll).Methods("GET")
	r.HandleFunc("/api/authors/{id}", author_api.Find).Methods("GET")
	r.HandleFunc("/api/authors/search/{keyword}", author_api.Search).Methods("GET")
	r.HandleFunc("/api/authors", author_api.Create).Methods("POST")
	r.HandleFunc("/api/authors/{id}", author_api.Update).Methods("PUT")
	r.HandleFunc("/api/authors", author_api.Save).Methods("PUT")
	r.HandleFunc("/api/authors/{id}", author_api.Delete).Methods("DELETE")

	//products
	r.HandleFunc("/api/products", product_api.FindAll).Methods("GET")
	r.HandleFunc("/api/products/search/{keyword}", product_api.Search).Methods("GET")
	r.HandleFunc("/api/products/{id}", product_api.Find).Methods("GET")
	r.HandleFunc("/api/products", product_api.Create).Methods("POST")
	r.HandleFunc("/api/products/{id}", product_api.Delete).Methods("DELETE")
	r.HandleFunc("/api/products", product_api.Save).Methods("PUT")
	r.HandleFunc("/api/products/{id}", product_api.Update).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}

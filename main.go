package main

import (
	"./src/apis/product_api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/products/findall", product_api.FindAll).Methods("GET")
	r.HandleFunc("/api/products/search/{keyword}", product_api.Search).Methods("GET")
	r.HandleFunc("/api/products/{id}", product_api.Find).Methods("GET")
	r.HandleFunc("/api/products", product_api.Create).Methods("POST")
	r.HandleFunc("/api/products/{id}", product_api.Delete).Methods("DELETE")
	r.HandleFunc("/api/products", product_api.Save).Methods("PUT")
	r.HandleFunc("/api/products/{id}", product_api.Update).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}

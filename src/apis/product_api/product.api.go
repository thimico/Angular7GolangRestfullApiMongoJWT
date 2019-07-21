package product_api

import (
	"../../models"
	"encoding/json"
	"../../entities"
	"github.com/gorilla/mux"
	"net/http"
)

func FindAll(res http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	products, err := productModel.FindAll()
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, products)
	}
}

func Search(res http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	products, err := productModel.Search(keyword)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {

		respondWithJson(res, http.StatusOK, products)
	}
}

func Find(res http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	vars := mux.Vars(request)
	id := vars["id"]
	product, err := productModel.Find(id)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {

		respondWithJson(res, http.StatusOK, product)
	}
}

func Create(res http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var productModel models.ProductModel
		err2 := productModel.Create(&product)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(res, http.StatusOK, product)
		}
	}
}

func Update(res http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		vars := mux.Vars(request)
		id := vars["id"]
		var productModel models.ProductModel
		err2 := productModel.Update(id, &product)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(res, http.StatusOK, product)
		}
	}
}

func Save(res http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var productModel models.ProductModel
		err2 := productModel.Save(&product)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err.Error())
		} else {
			respondWithJson(res, http.StatusOK, product)
		}
	}
}

func Delete(res http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	vars := mux.Vars(request)
	id := vars["id"]
	product, err := productModel.Find(id)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		err2 := productModel.Delete(product)
		if err2 != nil {
			respondWithError(res, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(res, http.StatusOK, product)
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

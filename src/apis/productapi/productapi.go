package apis

import (
	"encoding/json"
	"net/http"

	"../../config"
	"../../dao"
	a "../../dao/abstractdao"
	"../../entities"
	. "../abstractapi"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type ProductAPI struct {
	AbstractAPI
}

func (ext ProductAPI) ObterTodos(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {

		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "product"}
		productDAO := dao.ProductDAO{AbstractDAO: abstractDAO}
		produtos, err2 := productDAO.FindAll()

		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, produtos)
		}
	}

}

// Add new product
func (ext ProductAPI) Create(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "product"}
		productDAO := dao.ProductDAO{AbstractDAO: abstractDAO}

		var product entities.Product
		product.Id = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := productDAO.Create(&product)
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				respondWithJson(response, http.StatusOK, product)
			}
		}
	}
}

func (ext ProductAPI) Find(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "product"}
		productDAO := dao.ProductDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		product, err2 := productDAO.Find(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

func (ext ProductAPI) Delete(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "product"}
		productDAO := dao.ProductDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		err2 := productDAO.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, entities.Product{})
		}
	}
}

func (ext ProductAPI) Update(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "product"}
		productDAO := dao.ProductDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		var product entities.Product
		err2 := json.NewDecoder(request.Body).Decode(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := productDAO.Update(id, &product)
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				respondWithJson(response, http.StatusOK, product)
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

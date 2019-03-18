package accountapi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../../config"
	"../../dao"
	a "../../dao/abstractdao"
	"../../entities"
	"gopkg.in/mgo.v2/bson"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "account"}
		accountDAO := dao.AccountDAO{AbstractDAO: abstractDAO}
		account, err2 := accountDAO.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, account)
		}
	}
}

func Find(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "account"}
		accountDAO := dao.AccountDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		account, err2 := accountDAO.Find(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, account)
		}
	}
}

// Add new account
func Create(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "account"}
		accountDAO := dao.AccountDAO{AbstractDAO: abstractDAO}
		var account entities.Account
		account.Id = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&account)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := accountDAO.Create(&account)
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				respondWithJson(response, http.StatusOK, account)
			}
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "account"}
		accountDAO := dao.AccountDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		err2 := accountDAO.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, entities.Account{})
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "account"}
		accountDAO := dao.AccountDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		var account entities.Account
		err2 := json.NewDecoder(request.Body).Decode(&account)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := accountDAO.Update(id, &account)
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				respondWithJson(response, http.StatusOK, account)
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

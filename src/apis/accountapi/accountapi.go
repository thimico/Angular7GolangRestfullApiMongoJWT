package accountapi

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"../../config"
	"../../entities"
	"../../models"
)

func Create(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		accountModel := models.AccountModel{
			DB: db,
		}
		var account entities.Account
		account.Id = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&account)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := accountModel.Create(&account)
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

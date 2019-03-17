package jwtauth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"../../config"
	"../../entities"
	"../../models"
)

var secretKey = "MySecretKey"

func GenerateToken(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		respondWithError(response, http.StatusUnauthorized, err.Error())
	} else {
		db, err2 := config.Connect()
		if err2 != nil {
			respondWithError(response, http.StatusUnauthorized, err2.Error())
		} else {
			accountModel := models.AccountModel{
				DB: db,
			}
			valid := accountModel.CheckUsernameAndPassword(account.Username, account.Password)
			if valid {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"username": account.Username,
					"password": account.Password,
					"exp":      time.Now().Add(time.Hour * 72).Unix(),
				})
				tokenString, err2 := token.SignedString([]byte(secretKey))
				if err2 != nil {
					respondWithError(response, http.StatusUnauthorized, err.Error())
				} else {
					respondWithJson(response, http.StatusOK, entities.JWTToken{Token: tokenString})
				}
			} else {
				respondWithError(response, http.StatusUnauthorized, "Account Invalid")
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

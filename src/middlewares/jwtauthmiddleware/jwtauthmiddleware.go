package jwtauthmiddleware

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"../../config"
	"../../models"
)

var secretKey = "MySecretKey"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		stringToken := request.Header.Get("Authorization")
		if stringToken == "" {
			respondWithError(response, http.StatusUnauthorized, "Unauthorized")
		} else {
			result, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if err != nil {
				respondWithJson(response, http.StatusUnauthorized, err.Error())
				return
			} else {
				if result.Valid {
					next.ServeHTTP(response, request)
				} else {
					respondWithJson(response, http.StatusUnauthorized, "Invalid Authorization")
				}
			}
			next.ServeHTTP(response, request)
		}
	})
}

func CheckUsernameAndPassword(username, password string) bool {
	db, err := config.Connect()
	if err != nil {
		return false
	} else {
		accountModel := models.AccountModel{
			DB: db,
		}
		return accountModel.CheckUsernameAndPassword(username, password)
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

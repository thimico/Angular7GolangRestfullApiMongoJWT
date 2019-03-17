package basicauth

import (
	"net/http"

	"../../config"
	"../../models"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth()
		if !ok || !CheckUsernameAndPassword(username, password) {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Unauthorized"))
		} else {
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

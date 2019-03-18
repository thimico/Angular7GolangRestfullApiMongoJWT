package basicauth

import (
	"net/http"

	"../../config"
	"../../dao"
	a "../../dao/abstractdao"
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
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "account"}
		accountDAO := dao.AccountDAO{AbstractDAO: abstractDAO}
		return accountDAO.CheckUsernameAndPassword(username, password)
	}
}

func CheckEmailAndPassword(email, password string) bool {
	db, err := config.Connect()
	if err != nil {
		return false
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "account"}
		accountDAO := dao.AccountDAO{AbstractDAO: abstractDAO}
		return accountDAO.CheckEmailAndPassword(email, password)
	}
}

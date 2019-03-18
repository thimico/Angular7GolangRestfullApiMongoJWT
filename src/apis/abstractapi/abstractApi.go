package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	dao "../../dao/abstractdao"
	"gopkg.in/mgo.v2"
)

type IAPI interface {
	Hi(http.ResponseWriter, *http.Request)
	FindAll(http.ResponseWriter, *http.Request)
	respondWithError(http.ResponseWriter, int, string)
	respondWithJson(http.ResponseWriter, int, interface{})
}

type AbstractAPI struct {
	FirstName string
	// IAPI
}

func (g AbstractAPI) GetName() string {
	return g.FirstName
}

func (g AbstractAPI) OldPrint() {
	fmt.Printf("Hello, %s\n", g.GetName())
}

func (g AbstractAPI) Hi(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hi from Product")
}

func (g AbstractAPI) FindAll(db *mgo.Database, dao dao.AbstractDAO, response http.ResponseWriter, request *http.Request) {
	results, err2 := dao.FindAll()
	if err2 != nil {
		g.respondWithError(response, http.StatusBadRequest, err2.Error())
		return
	} else {
		g.respondWithJson(response, http.StatusOK, results)
	}
}

func (g AbstractAPI) respondWithError(w http.ResponseWriter, code int, msg string) {
	g.respondWithJson(w, code, map[string]string{"error": msg})
}

func (g AbstractAPI) respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

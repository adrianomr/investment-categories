package rest

import (
	"adrianorodrigues.com.br/investment-categories/framework/external/rest/dto"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var singleton = HttpServerImpl{}

type HttpServer interface {
	Init()
}

func HttpServerSingleton() *HttpServerImpl {
	return &singleton
}

type HttpServerImpl struct {
}

func (h HttpServerImpl) Init() {
	router := mux.NewRouter()
	router.HandleFunc("/categories", CategoriesControllerSingleton().GetCategory).Methods("GET")
	router.HandleFunc("/categories", CategoriesControllerSingleton().PostCategory).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func sendResponse(w http.ResponseWriter, response dto.ResponseDto) {
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func readRequest(r *http.Request, bodyResult interface{}) {
	json.NewDecoder(r.Body).Decode(&bodyResult)
}

package rest

import (
	"adrianorodrigues.com.br/investment-categories/framework/entrypoint/rest/dto"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
)

var singleton = HttpServerImpl{}

type HttpServer interface {
	Init()
}

func HttpServerSingleton() *HttpServerImpl {
	return &singleton
}

type HttpServerImpl struct {
	server *http.Server
}

func (h HttpServerImpl) Init() {
	router := h.buildRouter()
	h.server = &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	log.Fatal(h.server.ListenAndServe())
}

func (h HttpServerImpl) buildRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/categories", CategoriesControllerSingleton().GetCategory).Methods("GET")
	router.HandleFunc("/categories", CategoriesControllerSingleton().PostCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", CategoriesControllerSingleton().PutCategory).Methods("PUT")
	router.Use(commonMiddleware)
	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (h HttpServerImpl) InitTest(req *http.Request) *httptest.ResponseRecorder {
	router := h.buildRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func sendResponse(w http.ResponseWriter, response dto.ResponseDto) {
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func readRequest(r *http.Request, bodyResult interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&bodyResult)
	return err
}

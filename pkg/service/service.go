package service

import (
	"github.com/gorilla/mux"
	"net/http"
)

//API represents the API exposed by the service
type API struct {
	endpoints []Endpoint
	router    *mux.Router
}

//Endpoint represents a single API endpoint
type Endpoint struct {
	path    string
	method  string
	handler http.Handler
}

//NewAPI creates a new API instance
func NewAPI() *API {
	api := new(API)
	api.router = mux.NewRouter()
	return api
}

//AddEndpoint creates a new endpoint
func (api *API) AddEndpoint(path string, method string, handler http.Handler) {
	api.router.Handle(path, handler).Methods(method)
}

//Serve api in the specified port
func Serve(api *API, port string) error {
	return http.ListenAndServe(port, api.router)
}

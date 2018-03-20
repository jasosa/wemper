package main

import (
	"github.com/gorilla/mux"
	"github.com/jasosa/wemper/pkg/service"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/persons", service.APIErrorMiddleware(service.GetAllUsersHandler)).Methods("GET")
	router.Handle("/persons/{id}/invitations/", service.APIErrorMiddleware(service.InvitePersonHandler)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

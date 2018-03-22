package main

import (
	"github.com/gorilla/mux"
	"github.com/jasosa/wemper/pkg/service"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	handler := service.APIErrorMiddleware(service.GetAllUsersHandler)
	loggedHandler := service.LoggingMiddleware(handler)
	router.Handle("/persons", loggedHandler).Methods("GET")
	router.Handle("/persons/{id}/invitations/", service.APIErrorMiddleware(service.InvitePersonHandler)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

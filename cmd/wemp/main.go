package main

import (
	"github.com/gorilla/mux"
	"github.com/jasosa/wemper/pkg/service"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	getAllUsersHandlerWithHTTPError := service.HTTPErrorMiddleware(service.GetAllUsersHandler)
	invitePersonsHandlerWithHTTPError := service.HTTPErrorMiddleware(service.InvitePersonHandler)

	router.Handle("/persons",
		service.LoggingMiddleware(getAllUsersHandlerWithHTTPError)).Methods("GET")
	router.Handle("/persons/{id}/invitations/",
		service.LoggingMiddleware(invitePersonsHandlerWithHTTPError)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

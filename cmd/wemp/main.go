package main

import (
	"github.com/jasosa/wemper/pkg/server"
)

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/persons", server.AppHandler(server.GetAllUsers)).Methods("GET")
	router.Handle("/persons/{id}/invitations/", server.AppHandler(server.InvitePerson)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

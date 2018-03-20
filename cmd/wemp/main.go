package main

import (
	"github.com/jasosa/wemper/pkg/service"
)

import (
	"log"
)

func main() {
	/* router := mux.NewRouter()
	router.Handle("/persons", service.AppHandler(service.GetAllUsers)).Methods("GET")
	router.Handle("/persons/{id}/invitations/", service.AppHandler(service.InvitePerson)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router)) */

	api := service.NewAPI()
	api.AddEndpoint("/persons", "GET", service.APIErrorHandler(service.GetAllUsers))
	api.AddEndpoint("/persons/{id}/invitations/", "POST", service.APIErrorHandler(service.InvitePerson))
	log.Fatal(service.Serve(api, ":8080"))
}

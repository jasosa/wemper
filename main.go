package main

import (
	"feedwell/invitations"
	"feedwell/mysql"
	"feedwell/server"
)

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//GetAllUsers returns all users (registered and invited) from the community. We start with only 1 default community, so all people will be returned
func GetAllUsers(w http.ResponseWriter, r *http.Request) error {
	// TODO: Check errors in encoding
	users, err := service.GetAllUsers(repository.GetAllPeople)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(users)
}

// InvitePerson sends an invitation to the person to the default community. This person have the status "invited"
func InvitePerson(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	var newUser invitations.NewUser
	errDecode := json.NewDecoder(r.Body).Decode(&newUser)
	if errDecode != nil {
		return &server.RequestDecodeError{BaseError: errDecode}
	}

	inv, err := service.InvitePerson(id, newUser)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(inv)

}

var service invitations.Service
var repository invitations.Repository

func init() {
	// assigning all the repository interfaces
	//repository = inmemorydb.NewPeopleRepository()
	repository = mysql.NewPeopleRepository(new(mysql.DbConnection))
	var sender = invitations.NewFakeInvitationSender()
	service = invitations.NewBasicService(repository, sender)
}

func main() {
	router := mux.NewRouter()
	router.Handle("/people", server.AppHandler(GetAllUsers)).Methods("GET")
	router.Handle("/people/{id}/invitations/", server.AppHandler(InvitePerson)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

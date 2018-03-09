package main

import (
	"encoding/json"
	"feedwell/inmemorydb"
	"feedwell/people"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetAllUsers returns all users (registered and invited) from the community. We start with only 1 default community, so all people will be returned
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: Check errors in encoding
	json.NewEncoder(w).Encode(service.GetAllUsers())
}

// InvitePerson sends an invitation to the person to the default community. This person have the status "invited"
func InvitePerson(w http.ResponseWriter, r *http.Request) {
	//TODO: Validations creating new people
	// TODO: Check errors in decoding
	//params := mux.Vars(r)
	var person people.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	service.InvitePerson(person)
	//json.NewEncoder(w).Encode(repository.GetAllPeople())
}

var service people.Service

func init() {
	// assigning all the repository interfaces
	var repository = inmemorydb.NewPeopleRepository()
	service = people.NewBasicService(repository)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetAllUsers).Methods("GET")
	router.HandleFunc("/people/", InvitePerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

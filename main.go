package main

import (
	"encoding/json"
	"feedwell/inmemorydb"
	"feedwell/people"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetPeople returns all people from the community. We start with only 1 default community, so all people will be returned
func GetPeople(w http.ResponseWriter, r *http.Request) {
	// TODO: Check errors in encoding
	json.NewEncoder(w).Encode(repository.GetAllPeople())
}

// InvitePerson invites a person to the default community. This person  have the status "invited"
func InvitePerson(w http.ResponseWriter, r *http.Request) {
	//TODO: Validations creating new people
	// TODO: Check errors in decoding
	params := mux.Vars(r)
	var person people.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	repository.AddPerson(person)
	json.NewEncoder(w).Encode(repository.GetAllPeople())
}

var repository people.Repository

func init() {
	// assigning all the repository interfaces
	repository = inmemorydb.NewPeopleRepository()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", InvitePerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

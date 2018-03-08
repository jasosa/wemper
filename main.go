package main

import (
	"encoding/json"
	"feedwell/people"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetPeople returns all people from the community. We start with only 1 default community, so all people will be returned
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people.People)
}

// InvitePerson invites a person to the default community. This person  have the status "invited"
func InvitePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person people.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people.People = append(people.People, person)
	json.NewEncoder(w).Encode(people.People)
}

func main() {
	router := mux.NewRouter()
	//TODO: Move out from here population of people
	people.People = append(people.People, people.Person{ID: "1", Name: "John Doe", Email: "john@doe.com"})
	people.People = append(people.People, people.Person{ID: "2", Name: "John Brown", Email: "john@brown.com"})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", InvitePerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

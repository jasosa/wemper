package main

import (
	"encoding/json"
	"feedwell/fake"
	"feedwell/invitations"
	"feedwell/mysql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetAllUsers returns all users (registered and invited) from the community. We start with only 1 default community, so all people will be returned
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: Check errors in encoding
	//json.NewEncoder(w).Encode(service.GetAllUsers())
	users, _ := service.GetAllUsers(repository.GetAllPeople)
	json.NewEncoder(w).Encode(users)
}

// InvitePerson sends an invitation to the person to the default community. This person have the status "invited"
func InvitePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var newUser invitations.NewUser
	errDecode := json.NewDecoder(r.Body).Decode(&newUser)
	if errDecode != nil {
		json.NewEncoder(w).Encode(errDecode.Error())
	} else {
		inv, err := service.InvitePerson(id, newUser)
		if err != nil {
			json.NewEncoder(w).Encode(err) //TODO: Change this to handle http Errors
		} else {
			json.NewEncoder(w).Encode(inv)
		}
	}
}

var service invitations.Service
var repository invitations.Repository

func init() {
	// assigning all the repository interfaces
	//repository = inmemorydb.NewPeopleRepository()
	repository = mysql.NewPeopleRepository(new(mysql.DbConnection))
	var sender = fake.NewInvitationSender()
	service = invitations.NewBasicService(repository, sender)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetAllUsers).Methods("GET")
	router.HandleFunc("/people/{id}/invitations/", InvitePerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

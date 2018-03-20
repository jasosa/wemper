package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jasosa/wemper/pkg/invitations"
	"github.com/jasosa/wemper/pkg/invitations/mysql"
	"net/http"
)

var repository invitations.Repository
var service invitations.Service

func init() {
	// assigning all the repository interfaces
	repository = mysql.NewPeopleRepository(new(mysql.DbConnection))
	var sender = invitations.NewFakeInvitationSender()
	service = invitations.NewBasicService(repository, sender)
}

//AppHandler ...
type AppHandler func(http.ResponseWriter, *http.Request) error

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		httperror := getHTTPError(err)
		http.Error(w, httperror.ErrorMessage, httperror.HTTPStatusCode)
	}
}

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
		return &RequestDecodeError{BaseError: errDecode}
	}

	inv, err := service.InvitePerson(id, newUser)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(inv)
}

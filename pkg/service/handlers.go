package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jasosa/wemper/pkg/invitations"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//getAllUsersHandler returns all users (registered and invited) from the community. We start with only 1 default community, so all people will be returned
func (svc *Service) getAllUsersHandler(w http.ResponseWriter, r *http.Request) error {
	// TODO: Check errors in encoding

	users, err := svc.invitationsAPI.GetAllUsers(svc.invitationsSource.GetAllPeople)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(users)
}

// invitePersonHandler sends an invitation to the person to the default community. This person have the status "invited"
func (svc *Service) invitePersonHandler(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	var newUser invitations.NewUser
	errDecode := json.NewDecoder(r.Body).Decode(&newUser)
	if errDecode != nil {
		return &RequestDecodeError{BaseError: errDecode}
	}

	inv, err := svc.invitationsAPI.InvitePerson(id, newUser)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(inv)
}

//ErrorHandler wraps a handler that return errors
//and convert them in the proper HTTP Error
type ErrorHandler struct {
	handlerFunc HandleWithErrorFunc
	Logger      *log.Logger
}

//HandleWithErrorFunc gets any error from the decorated handler and converts it into an ClientAPIError
type HandleWithErrorFunc func(http.ResponseWriter, *http.Request) error

func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handlerFunc(w, r); err != nil {
		h.Logger.WithError(err).WithFields(log.Fields{
			"requestedUrl": r.RequestURI,
		}).Error("Error calling handler function")
		httperror := getHTTPErrorFromDomainError(err)
		httperror.Write(w)
	}
}

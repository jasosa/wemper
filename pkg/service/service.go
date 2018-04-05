package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jasosa/wemper/pkg/invitations"
	"github.com/jasosa/wemper/pkg/invitations/mysql"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//Service represents the invitations service
type Service struct {
	invitationsAPI    invitations.API
	invitationsSource invitations.Source
	invitationsSender invitations.Sender
	logger            *log.Logger
	//TODO: add other things here like authorization,etc...
}

//New creates a new instance of service
func New(conf Config) *Service {
	svc := Service{
		logger: conf.Logger,
	}
	svc.invitationsSender = invitations.NewFakeInvitationSender()
	svc.invitationsSource = mysql.NewPeopleSource(new(mysql.DbConnection), conf.DBUser, conf.DBPwd, conf.DBName, conf.DBHost)
	svc.invitationsAPI = invitations.NewAPI(svc.invitationsSource, svc.invitationsSender)
	return &svc
}

//Server returns a server with all endpoints setup
func (svc *Service) Server(port string) *http.Server {
	router := mux.NewRouter()

	getAllUsersHandler := ErrorHandler{HandleWithErrorFunc(svc.getAllUsersHandler), svc.logger}
	router.Handle("api/persons", LoggingMiddleware(svc.logger, getAllUsersHandler)).Methods("GET")

	invitePersonsHandler := ErrorHandler{HandleWithErrorFunc(svc.invitePersonHandler), svc.logger}
	router.Handle("api/persons/{id}/invitations/", LoggingMiddleware(svc.logger, invitePersonsHandler)).Methods("POST")

	return &http.Server{Addr: port, Handler: router}
}

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
		httperror := getClientAPIError(err)
		http.Error(w, httperror.ErrorMessage, httperror.HTTPStatusCode)
	}
}

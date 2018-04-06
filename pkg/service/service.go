package service

import (
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
	router.Handle("/api/persons", (LoggingMiddleware(svc.logger, getAllUsersHandler))).Methods("GET")

	invitePersonsHandler := ErrorHandler{HandleWithErrorFunc(svc.invitePersonHandler), svc.logger}
	router.Handle("/api/persons/{id}/invitations/", LoggingMiddleware(svc.logger, invitePersonsHandler)).Methods("POST")

	return &http.Server{Addr: port, Handler: router}
}

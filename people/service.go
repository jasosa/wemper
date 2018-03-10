package people

import (
	"feedwell/invitations"
)

//Service exposes all the methods related to handled people
type Service interface {
	GetAllUsers() []User
	InvitePerson(inviterID string, invitee Person) (invitations.Invitation, error)
	//AcceptInvitation(invited Person)
}

//basicService implements Service
type basicService struct {
	repository Repository
	sender     invitations.Sender
}

//NewBasicService creates a new instance of people.Service
func NewBasicService(repository Repository, sender invitations.Sender) Service {
	service := new(basicService)
	service.repository = repository
	service.sender = sender
	return *service
}

//GetAllUsers ....
func (serv basicService) GetAllUsers() []User {
	return serv.repository.GetAllPeople()
}

//InvitePerson ... returns the generated invitation text
func (serv basicService) InvitePerson(inviterID string, invitee Person) (invitations.Invitation, error) {
	var invitation invitations.Invitation
	var err error
	var inviter *User

	//get inviter
	inviter, err = serv.repository.GetPersonByID(inviterID)
	if err == nil {
		//add person to the source
		serv.repository.AddPerson(&invitee)
		invitation, err = inviter.generateInvitation(invitee)
		if err == nil {
			//then send email to the person
			sendInvitation(invitation, serv.sender)
		}
	}
	return invitation, err
}

func sendInvitation(invitation invitations.Invitation, sender invitations.Sender) {
	sender.Send(invitation)
}

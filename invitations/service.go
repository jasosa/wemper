package invitations

import (
	"errors"
	"fmt"
)

//Service exposes all the methods related to handled people
type Service interface {
	GetAllUsers(GetFromSource func() ([]AppUser, error)) ([]AppUser, error)
	InvitePerson(inviterID string, invited NewUser) (Invitation, error)
	//AcceptInvitation(invited Person)
}

//basicService implements Service
type basicService struct {
	repository Repository
	sender     Sender
}

//NewBasicService creates a new instance of people.Service
func NewBasicService(repository Repository, sender Sender) Service {
	service := new(basicService)
	service.repository = repository
	service.sender = sender
	return *service
}

func (serv basicService) GetAllUsers(GetFromAnySource func() ([]AppUser, error)) ([]AppUser, error) {
	return GetFromAnySource()
}

//InvitePerson ... returns the generated invitation text
func (serv basicService) InvitePerson(inviterID string, newUser NewUser) (Invitation, error) {
	//get inviter
	inviterUser, err := serv.repository.GetPersonByID(inviterID)

	if err != nil {
		return Invitation{}, fmt.Errorf("Error getting information from datasource: %s", err.Error())
	}
	inviter, ok := inviterUser.(Inviter)
	if ok {
		//creates a new invitedUser based on newUser info
		invited := NewInvitedUser("", newUser.Name, newUser.Email)
		invitation := inviter.GenerateInvitation(invited)
		//new person is added
		serv.repository.AddPerson(invited)
		//then send email to the person
		sendInvitation(*invitation, serv.sender)
		return *invitation, nil
	}
	return Invitation{}, errors.New("The user has not enough permission to invite people")
}

func sendInvitation(invitation Invitation, sender Sender) {
	sender.Send(invitation)
}

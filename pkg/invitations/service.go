package invitations

import (
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
	repository Source
	sender     Sender
}

//NewBasicService creates a new instance of people.Service
func NewBasicService(repository Source, sender Sender) Service {
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

	inviterUser, err := serv.repository.GetPersonByID(inviterID)
	if err != nil {
		return Invitation{}, &UserNotFoundError{UserID: inviterID, BaseError: err}
	}

	inviter, ok := inviterUser.(Inviter)
	if ok {
		//creates a new invitedUser based on newUser info
		invited := NewInvitedUser("", newUser.Name, newUser.Email)
		if !invited.GetPersonInfo().HasValidNameAndEmail() {
			return Invitation{}, &UserNotValidError{
				BaseError: fmt.Errorf("Provided user info is not valid. \"Name\":{\"%s\"} \"Email\":{\"%s\"}", newUser.Name, newUser.Email),
			}
		}

		invitation := inviter.GenerateInvitation(invited)
		err = serv.repository.AddPerson(invited)
		if err != nil {
			return Invitation{}, &UserCouldNotBeAddedError{BaseError: err}
		}
		//then send email to the person
		sendInvitation(*invitation, serv.sender)
		return *invitation, nil
	}
	return Invitation{}, &ActionNotAllowedToUserError{UserID: inviterID, Action: "Invite"}
}

func sendInvitation(invitation Invitation, sender Sender) {
	sender.Send(invitation)
}

package invitations

import (
	"fmt"
)

//API exposes all the methods related to handled people
type API interface {
	GetAllUsers(GetFromSource func() ([]AppUser, error)) ([]AppUser, error)
	InvitePerson(inviterID string, invited NewUser) (Invitation, error)
	//AcceptInvitation(invited Person)
}

//API implements Service interface
type apiImpl struct {
	repository Source
	sender     Sender
}

//NewAPI creates a new instance of invitations.API
func NewAPI(repository Source, sender Sender) API {
	api := new(apiImpl)
	api.repository = repository
	api.sender = sender
	return *api
}

//GetAllUsers
func (api apiImpl) GetAllUsers(GetFromAnySource func() ([]AppUser, error)) ([]AppUser, error) {
	return GetFromAnySource()
}

//InvitePerson ... returns the generated invitation text
func (api apiImpl) InvitePerson(inviterID string, newUser NewUser) (Invitation, error) {

	inviterUser, err := api.repository.GetPersonByID(inviterID)
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
		err = api.repository.AddPerson(invited)
		if err != nil {
			return Invitation{}, &UserCouldNotBeAddedError{BaseError: err}
		}
		//then send email to the person
		sendInvitation(*invitation, api.sender)
		return *invitation, nil
	}
	return Invitation{}, &ActionNotAllowedToUserError{UserID: inviterID, Action: "Invite"}
}

func sendInvitation(invitation Invitation, sender Sender) {
	sender.Send(invitation)
}

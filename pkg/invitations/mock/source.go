package mock

import (
	"errors"
	"github.com/jasosa/wemper/pkg/invitations"
)

//Source mock source
type Source struct{}

//GetAllPeople from the mock source
func (s Source) GetAllPeople() ([]invitations.AppUser, error) {

	ru := invitations.NewRegisteredUser("1", "myname", "myemail")
	ru2 := invitations.NewRegisteredUser("2", "myname2", "myemail2")

	return []invitations.AppUser{
		ru,
		ru2,
	}, nil
}

//GetPersonByID from the mock source
func (s Source) GetPersonByID(id string) (invitations.AppUser, error) {
	if id == "nonExistentPersonID" {
		return nil, errors.New("This error should be sent")
	} else if id == "NonRegisteredPersonID" {
		return invitations.NewInvitedUser("1", "invitedPersonName", "invitedPersonEmail"), nil
	} else {
		return invitations.NewAdminUser("1", "adminPersonName", "adminPersonEmail"), nil
	}
}

//AddPerson ...
func (s Source) AddPerson(p invitations.AppUser) error {
	if p.GetPersonInfo().Name == "personThatWillFailWhenAdded" {
		return errors.New("simulating error")
	}
	return nil
}

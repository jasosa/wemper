package invitations

import (
	"errors"
)

//RepoForTest ...
type RepoForTest struct{}

//GetAllPeople ...
func (r RepoForTest) GetAllPeople() ([]AppUser, error) {

	ru := NewRegisteredUser("1", "myname", "myemail")
	ru2 := NewRegisteredUser("2", "myname2", "myemail2")

	return []AppUser{
		ru,
		ru2,
	}, nil
}

//GetPersonByID ...
func (r RepoForTest) GetPersonByID(id string) (AppUser, error) {
	if id == "nonExistentPersonID" {
		return nil, errors.New("This error should be sent")
	} else if id == "NonRegisteredPersonID" {
		return NewInvitedUser("1", "invitedPersonName", "invitedPersonEmail"), nil
	} else {
		return NewAdminUser("1", "adminPersonName", "adminPersonEmail"), nil
	}
}

//AddPerson ...
func (r RepoForTest) AddPerson(p AppUser) error {
	if p.GetPersonInfo().Name == "personThatWillFailWhenAdded" {
		return errors.New("simulating error")
	}
	return nil
}

//InvitationSenderForTest ...
type InvitationSenderForTest struct {
	SendCalled bool
}

//Send ...
func (es *InvitationSenderForTest) Send(inv Invitation) error {
	es.SendCalled = true
	return nil
}

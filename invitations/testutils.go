package invitations

import (
	"errors"
)

//RepoForTest ...
type RepoForTest struct{}

//GetAllPeople ...
func (r RepoForTest) GetAllPeople() []User {
	return []User{
		User{
			PersonBase: Person{Email: "myemail", Name: "nyname", ID: "1"},
			Registered: true,
		},
		User{
			PersonBase: Person{Email: "myemail2", Name: "nyname2", ID: "2"},
			Registered: true,
		},
	}
}

//GetPersonByID ...
func (r RepoForTest) GetPersonByID(id string) (*User, error) {
	if id == "nonExistentPersonID" {
		return nil, errors.New("This error should be sent")
	} else if id == "NonRegisteredPersonID" {
		return &User{
			PersonBase: Person{Email: "myemail", Name: "nyname", ID: "1"},
			Registered: false,
		}, nil
	} else {
		return &User{
			PersonBase: Person{Email: "emailFromRegistered", Name: "NameFromRegistered", ID: "1"},
			Registered: true,
		}, nil
	}
}

//AddPerson ...
func (r RepoForTest) AddPerson(p *Person) {
	//don´t do nothing
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

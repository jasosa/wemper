package invitations

import (
	"errors"
	"fmt"
)

// Person represents a person in the system
type Person struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// User represents a person that belongs to a community. She could be registered
// or not in the system
type User struct {
	PersonBase Person
	Registered bool
}

//NewNonRegisteredUser sdfasdf
func NewNonRegisteredUser(p Person) User {
	user := new(User)
	user.PersonBase = p
	user.Registered = false
	return *user
}

// Inviter represents the ability to invite someone to join a certain community
type Inviter interface {
	generateInvitation(p Person) (string, error)
}

func (u User) generateInvitation(p Person) (Invitation, error) {
	var invitation Invitation
	var err error
	if u.Registered {
		invitation = *NewInvitation(u.PersonBase.ID, p.ID, fmt.Sprintf("This is an invitation from %s to %s", u.PersonBase.Name, p.Email))
	} else {
		err = errors.New("A non-registered user cannot send invitations")
	}
	return invitation, err
}

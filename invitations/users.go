package invitations

import (
	"fmt"
)

//AppUser is an interface to handle all the users in the app
type AppUser interface {
	CanInvite() bool
	CanGiveFeedback() bool
	GetPersonInfo() *Person
}

//Inviter represent the ability to send invitations
type Inviter interface {
	GenerateInvitation(p AppUser) *Invitation
}

// Person represents a person in the system in any possible status (invited, registered, admin)
type Person struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Registered bool   `json:"registered,omitempty"`
	Admin      bool   `json:"admin,omitempty"`
}

//NewUser represents information of someone that is not already in the system
type NewUser struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

//InvitedUser represents a user that has been invited to the system
type invitedUser struct {
	*Person
}

//CanInvite returns true if the user has the ability to invite other users, false otherwise
func (i invitedUser) CanInvite() bool {
	return false
}

//CanGiveFeedback returns true if the person can give or receive feedback
func (i invitedUser) CanGiveFeedback() bool {
	return false
}

//GetPersonInfo gets the underlying person
func (i invitedUser) GetPersonInfo() *Person {
	return i.Person
}

//NewInvitedUser ...
func NewInvitedUser(ID, name, email string) AppUser {
	user := new(invitedUser)
	user.Person = &Person{
		ID:         ID,
		Email:      email,
		Name:       name,
		Registered: false,
		Admin:      false,
	}
	return user
}

//RegisteredUser represents a user that has been already registered in the system
type registeredUser struct {
	*Person
}

//CanInvite returns true if the user has the ability to invite other users, false otherwise
func (r registeredUser) CanInvite() bool {
	return false
}

//CanGiveFeedback returns true if the person can give or receive feedback
func (r registeredUser) CanGiveFeedback() bool {
	return true
}

//GetPersonInfo gets the underlying person
func (r registeredUser) GetPersonInfo() *Person {
	return r.Person
}

//NewRegisteredUser ...
func NewRegisteredUser(ID, name, email string) AppUser {
	user := new(registeredUser)
	user.Person = &Person{
		ID:         ID,
		Email:      email,
		Name:       name,
		Registered: true,
		Admin:      false,
	}
	return user
}

//AdminUser represents a user that is administrator of the system
type adminUser struct {
	*Person
}

//CanInvite returns true if the user has the ability to invite other users, false otherwise
func (a adminUser) CanInvite() bool {
	return true
}

//CanGiveFeedback returns true if the person can give or receive feedback
func (a adminUser) CanGiveFeedback() bool {
	return true
}

//GetPersonInfo gets the underlying person
func (a adminUser) GetPersonInfo() *Person {
	return a.Person
}

func (a *adminUser) GenerateInvitation(u AppUser) *Invitation {
	var invitation Invitation
	invitation = *NewInvitation(a.ID, u.GetPersonInfo().ID, fmt.Sprintf("This is an invitation from %s to %s", a.Name, u.GetPersonInfo().Email))
	return &invitation
}

//NewAdminUser ...
func NewAdminUser(ID, name, email string) AppUser {
	user := new(adminUser)
	user.Person = &Person{
		ID:         ID,
		Email:      email,
		Name:       name,
		Registered: true,
		Admin:      true,
	}
	return user
}

package people

import (
	"testing"
)

func TestWhenNewNonRegisteredUserIsCreatedThenShouldNotBeRegistered(t *testing.T) {
	user := NewNonRegisteredUser(giveMeAPerson())
	if user.Registered {
		t.Errorf("User should not be registered")
	}
}
func TestRegisteredUserGeneratesInvitationCorrectly(t *testing.T) {

	u := User{
		PersonBase: giveMeAPerson(),
		Registered: true,
	}

	invitation, _ := u.generateInvitation(giveMeAnotherPerson())
	if invitation.Text != "This is an invitation from Peter to John@cool.com" {
		t.Errorf("Invitation was not generated successfully")
	}
}

func TestNonRegisteredUserCannotInvite(t *testing.T) {
	u := NewNonRegisteredUser(giveMeAPerson())
	_, err := u.generateInvitation(giveMeAnotherPerson())
	if err == nil {
		t.Errorf("Not registered users should not be able to invite anyone")
	}
}

func giveMeAPerson() Person {
	person := Person{ID: "1", Name: "Peter", Email: "Peter@cool.com"}
	return person
}

func giveMeAnotherPerson() Person {
	person := Person{ID: "2", Name: "John", Email: "John@cool.com"}
	return person
}

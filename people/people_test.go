package people

import (
	"testing"
)

func TestWhenNewNonRegisteredUserIsCreatedThenShouldNotBeRegistered(t *testing.T) {
	user := NewNonRegisteredUser(*giveMeAPerson())
	if user.Registered {
		t.Errorf("User should not be registered")
	}
}
func TestRegisteredUserGeneratesInvitationCorrectly(t *testing.T) {

	u := User{
		PersonBase: *giveMeAPerson(),
		Registered: true,
	}

	invitation, _ := u.invite("myadress@email.com")
	if expected := "This is an invitation from Peter to myadress@email.com"; expected != invitation {
		t.Errorf("Invitation sent was not the expected one: %s", invitation)
	}
}

func TestNonRegisteredUserCannotInvite(t *testing.T) {
	u := NewNonRegisteredUser(*giveMeAPerson())
	_, err := u.invite("invited@user.com")
	if err == nil {
		t.Errorf("Not registered users should not be able to invite anyone")
	}
}

func giveMeAPerson() *Person {
	person := Person{ID: "1", Name: "Peter", Email: "Peter@cool.com"}
	return &person
}

package invitations_test

import (
	"github.com/jasosa/wemper/invitations"
	"testing"
)

func TestWhenInvitedUserIsCreatedThenShouldNotBeAbleToInviteOrGiveFeedback(t *testing.T) {
	user := invitations.NewInvitedUser("", "name", "email")
	if user.CanInvite() || user.CanGiveFeedback() {
		t.Errorf("An invited user should not be able to invite or give feedback")
	}
}

func TestWhenRegisteredUserIsCreatedThenShouldNotBeAbleToInviteButToGiveFeedback(t *testing.T) {
	user := invitations.NewRegisteredUser("1", "name", "email")
	if user.CanInvite() {
		t.Errorf("A registered user should not be able to invite anyone")
	}

	if !user.CanGiveFeedback() {
		t.Errorf("A registered user should be able to give/receive feedback")
	}
}

func TestWhenAdminUserIsCreatedThenShouldBeAbleToDoEverything(t *testing.T) {
	user := invitations.NewAdminUser("1", "name", "email")
	if !user.CanInvite() {
		t.Errorf("An admin user should be able to invite everyone")
	}

	if !user.CanGiveFeedback() {
		t.Errorf("An admin user should be able to give/receive feedback")
	}
}

func TestAdminUserGeneratesInvitationCorrectly(t *testing.T) {
	user := invitations.NewAdminUser("1", "name1", "email1")
	userInvited := invitations.NewInvitedUser("", "name2", "email2")
	inv := user.(invitations.Inviter).GenerateInvitation(userInvited)
	if inv.Text != "This is an invitation from name1 to email2" {
		t.Errorf("Invitation not generated succesfully: %s", inv.Text)
	}
}

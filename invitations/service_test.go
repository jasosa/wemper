package invitations

import (
	"testing"
)

func TestWhenGetAllUsersThenRepositoryGetsTheUsers(t *testing.T) {
	testRepo := new(RepoForTest)
	s := NewBasicService(testRepo, nil)
	users, _ := s.GetAllUsers(testRepo.GetAllPeople)
	if len(users) != 2 {
		t.Errorf("Retrieving users from source was not successful. %d were returned instead of %d", len(users), 2)
	}
}

func TestWhenInvitationSentByNonExistentInviterShouldReturnError(t *testing.T) {
	testRepo := new(RepoForTest)
	s := NewBasicService(testRepo, nil)

	_, err := s.InvitePerson("nonExistentPersonID", NewUser{Name: "myname", Email: "myemail"})
	if err == nil {
		t.Error("An error should be returned when an invitation arrives from a non-existend ID")
	}
}

func TestWhenInvitationCannotbeSentShouldReturnError(t *testing.T) {
	testRepo := new(RepoForTest)
	s := NewBasicService(testRepo, nil)
	_, err := s.InvitePerson("NonRegisteredPersonID", NewUser{Name: "myname", Email: "myemail"})
	if err.Error() != "The user has not enough permission to invite people" {
		t.Error("An error should be returned when a non-registered user tries to generates an invitation")
	}
}

func TestWhenInvitationIsSentTheBodyOfTheInvitationIsReturned(t *testing.T) {
	testRepo := new(RepoForTest)
	sender := new(InvitationSenderForTest)
	s := NewBasicService(testRepo, sender)
	inv, _ := s.InvitePerson("ARegisteredPersonID", NewUser{Name: "myname", Email: "myemail"})
	if inv.Text != "This is an invitation from adminPersonName to myemail" {
		t.Errorf("Invitation was not returned succesfully: %s", inv.Text)
	}
	if !sender.SendCalled {
		t.Error("Invitation sender was not called")
	}
}

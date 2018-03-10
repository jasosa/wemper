package invitations

import (
	"testing"
)

func TestWhenGetAllUsersThenRepositoryGetsTheUsers(t *testing.T) {
	testRepo := new(RepoForTest)
	s := NewBasicService(testRepo, nil)
	users := s.GetAllUsers(testRepo.GetAllPeople)
	if len(users) != 2 {
		t.Errorf("Retrieving users from source was not successful. %d were returned instead of %d", len(users), 2)
	}
}

func TestWhenInvitationSentByNonExistentInviterShouldReturnError(t *testing.T) {
	testRepo := new(RepoForTest)
	s := NewBasicService(testRepo, nil)
	_, err := s.InvitePerson("nonExistentPersonID", Person{Email: "myemail", Name: "nyname", ID: "10"})
	if err == nil {
		t.Error("An error should be returned when an invitation arrives from a non-existend ID")
	}
}

func TestWhenInvitationCannotbeSentShouldReturnError(t *testing.T) {
	testRepo := new(RepoForTest)
	s := NewBasicService(testRepo, nil)
	_, err := s.InvitePerson("NonRegisteredPersonID", Person{Email: "myemail@email.com", Name: "nyname", ID: "10"})
	if err.Error() != "A non-registered user cannot send invitations" {
		t.Error("An error should be returned when a non-registered user tries to generates an invitation")
	}
}

func TestWhenInvitationIsSentTheBodyOfTheInvitationIsReturned(t *testing.T) {
	testRepo := new(RepoForTest)
	sender := new(InvitationSenderForTest)
	s := NewBasicService(testRepo, sender)
	inv, _ := s.InvitePerson("ARegisteredPersonID", Person{Email: "myemail@email.com", Name: "nyname", ID: "10"})
	if inv.Text != "This is an invitation from NameFromRegistered to myemail@email.com" {
		t.Error("Invitation was not returned succesfully")
	}
	if !sender.SendCalled {
		t.Error("Invitation sender was not called")
	}
}

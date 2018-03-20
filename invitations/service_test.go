package invitations_test

import (
	"github.com/jasosa/wemper/invitations"
	"github.com/jasosa/wemper/mock"
	"reflect"
	"testing"
)

func TestWhenGetAllUsersThenRepositoryGetsTheUsers(t *testing.T) {
	testRepo := new(mock.Source)
	s := invitations.NewBasicService(testRepo, nil)
	users, _ := s.GetAllUsers(testRepo.GetAllPeople)
	if len(users) != 2 {
		t.Errorf("Retrieving users from source was not successful. %d were returned instead of %d", len(users), 2)
	}
}

func TestWhenInvitationSentByNonExistentInviterShouldReturnError(t *testing.T) {
	testRepo := new(mock.Source)
	s := invitations.NewBasicService(testRepo, nil)

	_, err := s.InvitePerson("nonExistentPersonID", invitations.NewUser{Name: "myname", Email: "myemail"})

	if err == nil {
		t.Errorf("No error returned. Expected error of type {\"%s\"}", "*UserNotFoundError")
	} else {
		switch et := err.(type) {
		case *invitations.UserNotFoundError:
			//ok
		default:
			t.Errorf("Wrong error returned: Expected error of type {\"%s\"} but returned {\"%s\"}", "*UserNotFoundError", reflect.TypeOf(et))
		}
	}
}

func TestWhenNotValidUserNameOrEmailSentShouldReturnError(t *testing.T) {
	testRepo := new(mock.Source)
	s := invitations.NewBasicService(testRepo, nil)

	_, err := s.InvitePerson("ARegisteredPersonID", invitations.NewUser{Name: "", Email: ""})

	if err == nil {
		t.Errorf("No error returned. Expected error of type {\"%s\"}", "*UserNotValidError")
	} else {
		switch et := err.(type) {
		case *invitations.UserNotValidError:
			//ok
		default:
			t.Errorf("Wrong error returned: Expected error of type {\"%s\"} but returned {\"%s\"}", "*UserNotValidError", reflect.TypeOf(et))
		}
	}
}

func TestWhenInvitingPersonCanNotBeAddedToTheServiceShouldReturnError(t *testing.T) {
	testRepo := new(mock.Source)
	s := invitations.NewBasicService(testRepo, nil)

	_, err := s.InvitePerson("ARegisteredPersonID", invitations.NewUser{Name: "personThatWillFailWhenAdded", Email: "myemail"})

	if err == nil {
		t.Errorf("No error returned. Expected error of type {\"%s\"}", "*UserCouldNotBeAddedError")
	} else {
		switch et := err.(type) {
		case *invitations.UserCouldNotBeAddedError:
			//ok
		default:
			t.Errorf("Wrong error returned: Expected error of type {\"%s\"} but returned {\"%s\"}", "*UserCouldNotBeAddedError", reflect.TypeOf(et))
		}
	}
}

func TestWhenInvitationCannotbeSentShouldReturnError(t *testing.T) {
	testRepo := new(mock.Source)
	s := invitations.NewBasicService(testRepo, nil)
	_, err := s.InvitePerson("NonRegisteredPersonID", invitations.NewUser{Name: "myname", Email: "myemail"})

	if err == nil {
		t.Errorf("No error returned. Expected error of type {\"%s\"}", "*ActionNotAllowedToUserError")
	} else {
		switch et := err.(type) {
		case *invitations.ActionNotAllowedToUserError:
			//ok
		default:
			t.Errorf("Wrong error returned: Expected error of type {\"%s\"} but returned {\"%s\"}", "*ActionNotAllowedToUserError", reflect.TypeOf(et))
		}
	}
}

func TestWhenInvitationIsSentTheBodyOfTheInvitationIsReturned(t *testing.T) {
	testRepo := new(mock.Source)
	sender := new(mock.Sender)
	s := invitations.NewBasicService(testRepo, sender)
	inv, _ := s.InvitePerson("ARegisteredPersonID", invitations.NewUser{Name: "myname", Email: "myemail"})
	if inv.Text != "This is an invitation from adminPersonName to myemail" {
		t.Errorf("Invitation was not returned succesfully: %s", inv.Text)
	}
	if !sender.SendCalled {
		t.Error("Invitation sender was not called")
	}
}

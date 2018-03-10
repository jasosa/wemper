package inmemorydb

import (
	"feedwell/invitations"
	"testing"
)

func TestWhenGetAllPeopleReturns2Persons(t *testing.T) {
	peopleRep := NewPeopleRepository()
	p := peopleRep.GetAllPeople()
	if peopleExpected := 2; len(p) != peopleExpected {
		t.Errorf("GetAllPeople returned %d instead of %d", len(p), peopleExpected)
	}
}

func TestWhenAddANewPersonThenGetAllPeopleReturnsOneMorePerson(t *testing.T) {
	peopleRep := NewPeopleRepository()
	peopleRep.AddPerson(&invitations.Person{ID: "3", Name: "Another person", Email: "another@person.com"})
	p := peopleRep.GetAllPeople()
	if peopleExpected := 3; len(p) != peopleExpected {
		t.Errorf("After adding a new people, GetNewPeople returned %d instead of %d", len(p), peopleExpected)
	}
}

func TestWhenAddANewPersonIDOfPersonShouldBeSet(t *testing.T) {
	peopleRep := NewPeopleRepository()
	peopleRep.AddPerson(giveMeAPerson())
	p := peopleRep.GetAllPeople()
	if expectedID := "3"; expectedID != p[len(p)-1].PersonBase.ID {
		t.Errorf("After adding a new people the id was not sucessfully set")
	}
}

func TestGetPersonByIDShouldReturnsTheRightPerson(t *testing.T) {
	peopleRep := NewPeopleRepository()
	p, _ := peopleRep.GetPersonByID("2")
	if expectedID := "2"; expectedID != p.PersonBase.ID {
		t.Errorf("Get person by ID should return %s and it´s returning %s", expectedID, p.PersonBase.ID)
	}
}

func TestWhenGetPersonDoesNotFindPersonShouldReturnsError(t *testing.T) {
	peopleRep := NewPeopleRepository()
	_, err := peopleRep.GetPersonByID("54")
	if err == nil {
		t.Errorf("Get person by non-existend ID should return an error")
	}
}

func giveMeAPerson() *invitations.Person {
	person := invitations.Person{Name: "Peter", Email: "Peter@cool.com"}
	return &person
}

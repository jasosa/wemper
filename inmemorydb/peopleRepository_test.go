package inmemorydb

import (
	"feedwell/people"
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
	peopleRep.AddPerson(people.Person{ID: "3", Name: "Another person", Email: "another@person.com"})
	p := peopleRep.GetAllPeople()
	if peopleExpected := 3; len(p) != peopleExpected {
		t.Errorf("After adding a new people, GetNewPeople returned %d instead of %d", len(p), peopleExpected)
	}
}

func TestWhenAddANewPersonIDOfPersonShouldBeSet(t *testing.T) {
	peopleRep := NewPeopleRepository()
	peopleRep.AddPerson(*giveMeAPerson())
	p := peopleRep.GetAllPeople()
	if expectedID := "3"; expectedID != p[len(p)-1].PersonBase.ID {
		t.Errorf("After adding a new people the id was not sucessfully set")
	}
}

func giveMeAPerson() *people.Person {
	person := people.Person{Name: "Peter", Email: "Peter@cool.com"}
	return &person
}

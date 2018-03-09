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

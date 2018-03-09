package inmemorydb

import (
	"feedwell/people"
)

//PeopleRepository is a repository  for get/set people
type PeopleRepository struct {
	myPeople []people.Person
}

//GetAllPeople from a repository in memory
func (imr PeopleRepository) GetAllPeople() []people.Person {
	return imr.myPeople
}

//AddPerson adds a new person to the source repository
func (imr *PeopleRepository) AddPerson(p people.Person) {
	imr.myPeople = append(imr.myPeople, p)
}

//NewPeopleRepository initializes a in memory repository
func NewPeopleRepository() *PeopleRepository {
	imr := new(PeopleRepository)
	imr.myPeople = append(imr.myPeople, people.Person{ID: "1", Name: "John Doe", Email: "john@doe.com"})
	imr.myPeople = append(imr.myPeople, people.Person{ID: "2", Name: "John Brown", Email: "john@brown.com"})
	return imr
}

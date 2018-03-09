package inmemorydb

import (
	"feedwell/people"
	"strconv"
)

//PeopleRepository is a repository  for get/set people
type PeopleRepository struct {
	users []people.User
}

//GetAllPeople from a repository in memory
func (imr PeopleRepository) GetAllPeople() []people.User {
	return imr.users
}

//AddPerson adds a new person to the source repository
func (imr *PeopleRepository) AddPerson(p people.Person) {
	//TODO: Change this mechanism
	p.ID = strconv.Itoa(len(imr.users) + 1)
	imr.users = append(imr.users, *people.NewNonRegisteredUser(p))
}

//NewPeopleRepository initializes a in memory repository
func NewPeopleRepository() *PeopleRepository {
	imr := new(PeopleRepository)
	newUser1 := people.NewNonRegisteredUser(people.Person{ID: "1", Name: "John Doe", Email: "john@doe.com"})
	newUser2 := people.NewNonRegisteredUser(people.Person{ID: "2", Name: "John Brown", Email: "john@brown.com"})
	imr.users = append(imr.users, *newUser1)
	imr.users = append(imr.users, *newUser2)
	return imr
}

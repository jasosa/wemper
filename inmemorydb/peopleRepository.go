package inmemorydb

import (
	"errors"
	"feedwell/invitations"
	"strconv"
)

//PeopleRepository is a repository  for get/set people
type PeopleRepository struct {
	users []invitations.User
}

//GetAllPeople from a repository in memory
func (imr PeopleRepository) GetAllPeople() ([]invitations.User, error) {
	return imr.users, nil
}

//GetPersonByID ...
func (imr PeopleRepository) GetPersonByID(id string) (*invitations.User, error) {
	for i := range imr.users {
		if imr.users[i].PersonBase.ID == id {
			return &imr.users[i], nil
		}
	}
	return nil, errors.New("User does not exists in DB")
}

//AddPerson adds a new person to the source repository
func (imr *PeopleRepository) AddPerson(p *invitations.Person) error {
	//TODO: Change this mechanism
	p.ID = strconv.Itoa(len(imr.users) + 1)
	imr.users = append(imr.users, invitations.NewNonRegisteredUser(*p))
	return nil
}

//NewPeopleRepository initializes a in memory repository
func NewPeopleRepository() invitations.Repository {
	imr := new(PeopleRepository)
	newUser1 := invitations.NewNonRegisteredUser(invitations.Person{ID: "1", Name: "John Doe", Email: "john@doe.com"})
	newUser1.Registered = true
	newUser2 := invitations.NewNonRegisteredUser(invitations.Person{ID: "2", Name: "John Brown", Email: "john@brown.com"})
	imr.users = append(imr.users, newUser1)
	imr.users = append(imr.users, newUser2)
	return imr
}

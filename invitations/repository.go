package invitations

//Repository for people
type Repository interface {
	GetAllPeople() []User
	GetPersonByID(id string) (*User, error)
	AddPerson(p *Person)
}

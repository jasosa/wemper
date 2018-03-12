package invitations

//Repository for people
type Repository interface {
	GetAllPeople() ([]User, error)
	GetPersonByID(id string) (*User, error)
	AddPerson(p *Person) error
}

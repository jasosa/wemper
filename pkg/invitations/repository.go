package invitations

//Repository for people
type Repository interface {
	GetAllPeople() ([]AppUser, error)
	GetPersonByID(id string) (AppUser, error)
	AddPerson(p AppUser) error
}

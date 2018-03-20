package invitations

//Source for people
type Source interface {
	GetAllPeople() ([]AppUser, error)
	GetPersonByID(id string) (AppUser, error)
	AddPerson(p AppUser) error
}

package people

//Repository for people
type Repository interface {
	GetAllPeople() []Person
	AddPerson(p Person)
}

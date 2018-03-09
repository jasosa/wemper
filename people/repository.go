package people

//Repository for people
type Repository interface {
	GetAllPeople() []User
	AddPerson(p Person)
}

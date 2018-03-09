package people

//Service exposes all the methods related to handled people
type Service interface {
	GetAllUsers() []User
	InvitePerson(p Person)
	//AcceptInvitation(invited Person)
}

//basicService implements Service
type basicService struct {
	repository Repository
}

//NewBasicService creates a new instance of people.Service
func NewBasicService(repository Repository) Service {
	service := new(basicService)
	service.repository = repository
	return *service
}

//GetAllUsers ....
func (serv basicService) GetAllUsers() []User {
	return serv.repository.GetAllPeople()
}

//InvitePerson ...
func (serv basicService) InvitePerson(p Person) {
	//add person to the source
	serv.repository.AddPerson(p)
	//then send email to the person
}

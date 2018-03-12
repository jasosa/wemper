package invitations

//Service exposes all the methods related to handled people
type Service interface {
	GetAllUsers(GetFromSource func() ([]User, error)) ([]User, error)
	InvitePerson(inviterID string, invitee Person) (Invitation, error)
	//AcceptInvitation(invited Person)
}

//basicService implements Service
type basicService struct {
	repository Repository
	sender     Sender
}

//NewBasicService creates a new instance of people.Service
func NewBasicService(repository Repository, sender Sender) Service {
	service := new(basicService)
	service.repository = repository
	service.sender = sender
	return *service
}

func (serv basicService) GetAllUsers(GetFromAnySource func() ([]User, error)) ([]User, error) {
	return GetFromAnySource()
}

//InvitePerson ... returns the generated invitation text
func (serv basicService) InvitePerson(inviterID string, invitee Person) (Invitation, error) {
	var invitation Invitation
	var err error
	var inviter *User

	//get inviter
	inviter, err = serv.repository.GetPersonByID(inviterID)
	if err == nil {
		//add person to the source
		invitation, err = inviter.generateInvitation(invitee)
		if err == nil {
			//new person is added
			serv.repository.AddPerson(&invitee)
			//then send email to the person
			sendInvitation(invitation, serv.sender)
		}
	}
	return invitation, err
}

func sendInvitation(invitation Invitation, sender Sender) {
	sender.Send(invitation)
}

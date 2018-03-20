package invitations

//Sender allows sending invitations
type Sender interface {
	Send(inv Invitation) error
}

//FakeInvitationSender represents a sender that does nothing with the invitation
type FakeInvitationSender struct{}

//NewFakeInvitationSender creates an instance of an fake invitation sender
func NewFakeInvitationSender() Sender {
	sndr := new(FakeInvitationSender)
	return sndr
}

//Send ....
func (es *FakeInvitationSender) Send(inv Invitation) error {
	return nil
}

package fake

import (
	"feedwell/invitations"
)

//InvitationSender represents a sender that does nothing with the invitation
type InvitationSender struct{}

//NewInvitationSender creates an instance of an fake invitation sender
func NewInvitationSender() invitations.Sender {
	sndr := new(InvitationSender)
	return sndr
}

//Send ....
func (es *InvitationSender) Send(inv invitations.Invitation) error {
	return nil
}

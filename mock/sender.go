package mock

import (
	"github.com/jasosa/wemper/invitations"
)

//Sender mock
type Sender struct {
	SendCalled bool
}

//Send mock
func (es *Sender) Send(inv invitations.Invitation) error {
	es.SendCalled = true
	return nil
}

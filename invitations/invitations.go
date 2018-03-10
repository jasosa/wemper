package invitations

//Invitation represents the action of invite someone
type Invitation struct {
	FromID string `json:"fromid,omitempty"`
	ToID   string `json:"toid,omitempty"`
	Text   string `json:"text,omitempty"`
}

//NewInvitation creates a new instance of invitation
func NewInvitation(fromID string, toID string, text string) *Invitation {
	invitation := new(Invitation)
	invitation.FromID = fromID
	invitation.ToID = toID
	invitation.Text = text
	return invitation
}

//Sender allows sending invitations
type Sender interface {
	Send(inv Invitation) error
}

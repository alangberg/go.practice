package domain

type DirectMessage struct {
	From    string
	To      string
	Content string
	Seen    bool
}

func NewDirectMessage(userFrom string, userTo string, content string) *DirectMessage {

	directMessage := DirectMessage{
		userFrom,
		userTo,
		content,
		false,
	}

	return &directMessage
}

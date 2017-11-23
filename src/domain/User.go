package domain

type User struct {
	Username string
	Inbox    []*DirectMessage
}

func NewUser(username string) *User {

	inbox := make([]*DirectMessage, 0)

	user := User{
		username,
		inbox,
	}
	return &user
}

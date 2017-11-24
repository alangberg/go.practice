package domain

import (
	"time"
)

type Tweet struct {
	User *User
	Text string
	Date *time.Time
	Id   int
}

func NewTweet(user *User, text string) *Tweet {

	date := time.Now()

	tweet := Tweet{
		user,
		text,
		&date,
		0,
	}

	return &tweet
}

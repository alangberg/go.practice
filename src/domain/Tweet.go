package domain

import (
	"time"
)

type Tweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

func NewTweet(user string, text string) *Tweet {

	date := time.Now()

	tweet := Tweet{
		user,
		text,
		&date,
		0,
	}

	return &tweet
}

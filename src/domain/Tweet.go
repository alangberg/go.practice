package domain

import (
	"time"
)

type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet {

	date := time.Now()

	tweet := Tweet{
		user,
		text,
		&date,
	}

	return &tweet
}

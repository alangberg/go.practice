package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	GetUser()
	GetText()
	GetId()
	SetId()
	PrintableTweet()
}

type TextTweet struct {
	user *User
	text string
	date *time.Time
	id   int
}

type ImageTweet struct {
	TextTweet
	url string
}

type QuoteTweet struct {
	TextTweet
	quote *Tweet
}

func NewTextTweet(user *User, text string) *TextTweet {

	date := time.Now()

	tweet := TextTweet{
		user,
		text,
		&date,
		0,
	}

	return &tweet
}

func NewImageTweet(user *User, text, url string) *ImageTweet {
	date := time.Now()
	textTweet := NewTextTweet(user, text)

	tweet := ImageTweet{
		*textTweet,
		url,
	}

	return &tweet
}

func NewQuoteTweet(user *User, text string, quote *Tweet) *QuoteTweet {
	date := time.Now()
	textTweet := NewTextTweet(user, text)

	tweet := QuoteTweet{
		*textTweet,
		quote,
	}

	return &tweet
}

func (t *Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.User.Username, t.Text)
}

func (t *Tweet) String() string {
	return t.PrintableTweet()
}

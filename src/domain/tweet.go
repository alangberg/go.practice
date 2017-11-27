package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	GetUser() *User
	GetText() string
	GetId() int
	GetDate() *time.Time
	SetId(id int)
	PrintableTweet() string
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
	quote Tweet
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

	textTweet := NewTextTweet(user, text)

	tweet := ImageTweet{
		*textTweet,
		url,
	}

	return &tweet
}

func NewQuoteTweet(user *User, text string, quote Tweet) *QuoteTweet {

	textTweet := NewTextTweet(user, text)

	tweet := QuoteTweet{
		*textTweet,
		quote,
	}

	return &tweet
}

//TextTweet Methods

func (t *TextTweet) GetUser() *User {
	return t.user
}

func (t *TextTweet) GetText() string {
	return t.text
}
func (t *TextTweet) GetId() int {
	return t.id
}

func (t *TextTweet) GetDate() *time.Time {
	return t.date
}

func (t *TextTweet) SetId(newId int) {
	t.id = newId
}

func (t *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.user.Username, t.text)
}

func (t *TextTweet) String() string {
	return t.PrintableTweet()
}

//ImageTweet Methods

func (t *ImageTweet) GetUser() *User {
	return t.user
}

func (t *ImageTweet) GetText() string {
	return t.text
}
func (t *ImageTweet) GetId() int {
	return t.id
}

func (t *ImageTweet) GetDate() *time.Time {
	return t.date
}

func (t *ImageTweet) SetId(newId int) {
	t.id = newId
}

func (t *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \n %s", t.user.Username, t.text, t.url)
}

func (t *ImageTweet) String() string {
	return t.PrintableTweet()
}

//QuoteTweet Methods

func (t *QuoteTweet) GetUser() *User {
	return t.user
}

func (t *QuoteTweet) GetText() string {
	return t.text
}
func (t *QuoteTweet) GetId() int {
	return t.id
}

func (t *QuoteTweet) GetDate() *time.Time {
	return t.date
}

func (t *QuoteTweet) SetId(newId int) {
	t.id = newId
}

func (t *QuoteTweet) PrintableTweet() string {
	quote := (t.quote).PrintableTweet()
	return fmt.Sprintf("@%s: %s \n Quote: '%s'", t.user.Username, t.text, quote)
}

func (t *QuoteTweet) String() string {
	return t.PrintableTweet()
}

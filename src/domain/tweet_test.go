package domain_test

import (
	"fmt"
	"testing"

	"github.com/alangberg/go.tuiter/src/domain"
)

var defUser *domain.User
var defTweetText string
var defTweetImg string
var defTextTweet *domain.Tweet

func defaultUser() *domain.User {
	return domain.NewUser("defaultUser")
}

func defaultTweetText() string {
	return "Default tweet text"
}

func defaultTweetImg() string {
	return "http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
}

func defaultTextTweet() *domain.Tweet {
	return domain.NewTextTweet(defaultUser(), defaultTweetText())
}
func TestMain(m *testing.M) {
	defUser = defaultUser()
	defTweetText = defaultTweetText()
	defTweetImg = defaultTweetImg()

	defTextTweet = defaultTextTweet()
	m.Run()
}

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet(defUser, defTweetText)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := fmt.Sprintf("@%s: %s", defUser.Username, defTweetText)
	if text != expectedText {
		t.Errorf("The expected text is '%s' but was '%s'", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet(defUser, defTweetText, defTweetImg)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := fmt.Sprintf("@%s: %s %s", defUser, defTweetText, defaultTweetImg)
	if text != expectedText {
		t.Errorf("The expected text is '%s' but was '%s'", expectedText, text)
	}

}

/*
func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
*/

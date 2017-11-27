package domain_test

import (
	"fmt"
	"testing"

	"github.com/alangberg/go.tuiter/src/domain"
)

var defUser *domain.User
var defTweetText string
var defTweetImg string
var defTweetQuote *domain.TextTweet

var defTextTweet *domain.TextTweet
var defImageTweet *domain.ImageTweet
var defQuotedTweet *domain.QuoteTweet

func defaultUser() *domain.User {
	return domain.NewUser("defaultUser")
}

func defaultTweetText() string {
	return "Default tweet text"
}

func defaultTweetImage() string {
	return "http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
}

func defaultTweetQuote() *domain.TextTweet {
	quotedUser := domain.NewUser("quotedUser")
	return domain.NewTextTweet(quotedUser, "quoted text")
}

//tweets default
func defaultTextTweet() *domain.TextTweet {
	return domain.NewTextTweet(defaultUser(), defaultTweetText())
}

func defaultImageTweet() *domain.ImageTweet {
	return domain.NewImageTweet(defaultUser(), defaultTweetText(), defaultTweetImage())
}

func defaultQuotedTweet() *domain.QuoteTweet {
	return domain.NewQuoteTweet(defaultUser(), defaultTweetText(), defaultTweetQuote())
}

func TestMain(m *testing.M) {
	defUser = defaultUser()
	defTweetText = defaultTweetText()
	defTweetImg = defaultTweetImage()
	defTweetQuote = defaultTweetQuote()

	defTextTweet = defaultTextTweet()
	defImageTweet = defaultImageTweet()
	defQuotedTweet = defaultQuotedTweet()
	m.Run()
}

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := defTextTweet

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
	tweet := defImageTweet

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := fmt.Sprintf("@%s: %s \n %s", defUser.Username, defTweetText, defTweetImg)
	if text != expectedText {
		t.Errorf("The expected text is '%s' but was '%s'", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	tweet := defQuotedTweet

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedQuoteText := defTweetQuote.PrintableTweet()
	expectedText := fmt.Sprintf("@%s: %s \n Quote: '%s'", defUser.Username, defTweetText, expectedQuoteText)
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

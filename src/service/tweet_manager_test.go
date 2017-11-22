package service_test

import (
	"testing"

	"github.com/alangberg/go.tuiter/src/domain"
	"github.com/alangberg/go.tuiter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	user := "grupoEsfera"
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()

	if publishedTweet.User != user || publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \n but is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestCleanTweetDeletesTweet(t *testing.T) {
	user := "grupoEsfera"
	text := "Tweet to be deleted"

	tweet := domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	service.DeleteTweet()

	if service.GetTweet() != nil {
		t.Error("Expected tweet is '' ")
	}

}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var user string
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}

}

func TestCanNotPublishTweetWithoutText(t *testing.T) {
	var text string
	user := "grupoEsfera"
	tweet := domain.NewTweet(user, text)

	err := service.PublishTweet(tweet)

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestCanNotPublishTweetsLongerThan140Characters(t *testing.T) {
	text := `Go became a public open source project on November 10, 2009. After a couple of years of very active design and development, stability was called for and Go 1 was released on March 28, 2012. Go 1, which includes a language specification, standard libraries, and custom tools, provides a stable foundation for creating reliable products, projects, and publications.
	With that stability established, we are using Go to develop programs, products, and tools rather than actively changing the language and libraries. In fact, the purpose of Go 1 is to provide long-term stability. Backwards-incompatible changes will not be made to any Go 1 point release. We want to use what we have to learn how a future version of Go might look, rather than to play with the language underfoot. `
	user := "grupoEsfera"

	tweet := domain.NewTweet(user, text)

	err := service.PublishTweet(tweet)

	if err == nil || err.Error() != "text can not be longer than 140 characters" {
		t.Error("Expected error is text can not be longer than 140 characters")
	}

}

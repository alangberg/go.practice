package service_test

import (
	"testing"

	"github.com/go.tuiter/src/domain"

	"github.com/go.tuiter/src/service"
)

func GetDefaultTweet() *domain.Tweet {
	user := "grupoEsfera"
	text := "This is my first tweet"

	tweet := domain.NewTweet(user, text)
	return tweet
}

func TestPublishedTweetIsSaved(t *testing.T) {
	tweet := GetDefaultTweet()

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

	if service.GetTweet() != "" {
		t.Error("Expected tweet is '' ")
	}

}

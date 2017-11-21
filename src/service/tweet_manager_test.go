package service_test

import (
	"testing"

	"github.com/go.tuiter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	tweet := "This is my first tweet"
	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}

func TestCleanTweetDeletesTweet(t *testing.T) {
	tweet := "Tweet to be deleted"
	service.PublishTweet(tweet)

	service.DeleteTweet()

	if service.GetTweet() != "" {
		t.Error("Expected tweet is '' ")
	}

}

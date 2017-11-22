package service

import (
	"fmt"

	"github.com/alangberg/go.tuiter/src/domain"
)

var tweet *domain.Tweet

func PublishTweet(newTweet *domain.Tweet) error {

	if newTweet.User == "" {
		return fmt.Errorf("user is required")
	}

	tweet = newTweet
	return nil
}

func GetTweet() *domain.Tweet {
	return tweet
}

func DeleteTweet() {
	tweet = nil
}

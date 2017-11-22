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
	if newTweet.Text == "" {
		return fmt.Errorf("text is required")
	} else if len(newTweet.Text) > 140 {
		return fmt.Errorf("text can not be longer than 140 characters")
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

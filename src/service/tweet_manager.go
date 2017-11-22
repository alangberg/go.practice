package service

import (
	"github.com/alangberg/go.tuiter/src/domain"
)

var tweet *domain.Tweet

func PublishTweet(newTweet *domain.Tweet) {
	tweet = newTweet
}

func GetTweet() *domain.Tweet {
	return tweet
}

func DeleteTweet() {
	tweet = nil
}

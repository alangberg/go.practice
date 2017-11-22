package service

import (
	"fmt"

	"github.com/alangberg/go.tuiter/src/domain"
)

var tweets []*domain.Tweet

func PublishTweet(newTweet *domain.Tweet) (int, error) {
	if newTweet.User == "" {
		return -1, fmt.Errorf("user is required")
	}
	if newTweet.Text == "" {
		return -1, fmt.Errorf("text is required")
	} else if len(newTweet.Text) > 140 {
		return -1, fmt.Errorf("text can not be longer than 140 characters")
	}

	tweets = append(tweets, newTweet)
	tweetId := TweetCount() - 1
	newTweet.Id = tweetId

	return tweetId, nil
}

func GetTweetById(id int) *domain.Tweet {
	return tweets[id]
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func TweetCount() int {
	return len(tweets)
}

func DeleteTweets() {
	tweets = make([]*domain.Tweet, 0)
}

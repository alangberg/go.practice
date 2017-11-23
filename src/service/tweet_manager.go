package service

import (
	"fmt"

	"github.com/alangberg/go.tuiter/src/domain"
)

var tweetsMap map[string][]*domain.Tweet
var nextId int

func InitializeService() {
	tweetsMap = make(map[string][]*domain.Tweet)
	nextId = 0
}

func PublishTweet(newTweet *domain.Tweet) (int, error) {
	if newTweet.User == "" {
		return -1, fmt.Errorf("user is required")
	}
	if newTweet.Text == "" {
		return -1, fmt.Errorf("text is required")
	} else if len(newTweet.Text) > 140 {
		return -1, fmt.Errorf("text can not be longer than 140 characters")
	}

	tweets, ok := tweetsMap[newTweet.User]
	if !ok {
		tweetsMap[newTweet.User] = make([]*domain.Tweet, 0)
		tweets = tweetsMap[newTweet.User]
	}
	tweets = append(tweets, newTweet)
	tweetsMap[newTweet.User] = tweets
	newTweet.Id = nextId
	nextId++
	return newTweet.Id, nil
}

func GetTweetById(id int) *domain.Tweet {
	for _, tweets := range tweetsMap {
		for _, tweet := range tweets {
			if tweet.Id == id {
				return tweet
			}
		}
	}
	return nil
}

func GetTweetsByUser(user string) []*domain.Tweet {
	tweets, ok := tweetsMap[user]
	if ok {
		return tweets
	}
	return nil
}

func GetTweet() *domain.Tweet {
	return GetTweetById(nextId - 1)
}

func CountTweetsByUser(user string) int {
	tweets, ok := tweetsMap[user]
	if ok {
		return len(tweets)
	}
	return -1
}

func TweetCount() int {
	return nextId
}

func DeleteTweets() {
	InitializeService()
}

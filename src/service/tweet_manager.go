package service

import (
	"fmt"
	"sort"

	"github.com/alangberg/go.tuiter/src/domain"
)

var tweetsMap map[string][]*domain.Tweet
var followersMap map[string][]string
var nextId int

func InitializeService() {
	tweetsMap = make(map[string][]*domain.Tweet)
	followersMap = make(map[string][]string)
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
	tweetsMap = make(map[string][]*domain.Tweet)
	nextId = 0
}

func Follow(followingUser, newFollowedUser string) error {
	if CountTweetsByUser(newFollowedUser) == -1 {
		return fmt.Errorf("both users must exist")
	}

	followed, ok := followersMap[followingUser]

	if !ok {
		followersMap[followingUser] = make([]string, 0)
		followed = followersMap[followingUser]
	}

	followersMap[followingUser] = append(followed, newFollowedUser)

	return nil
}

func GetTimeline(user string) []*domain.Tweet {
	timeline := make([]*domain.Tweet, 0)
	followedUsers, ok := followersMap[user]

	if !ok {
		return nil
	}

	for _, followedUser := range followedUsers {
		timeline = append(timeline, tweetsMap[followedUser]...)
	}
	sort.Slice(timeline, func(i, j int) bool { return timeline[i].Date.Before(*timeline[j].Date) })
	return timeline
}

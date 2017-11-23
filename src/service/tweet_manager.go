package service

import (
	"fmt"
	"sort"

	"github.com/alangberg/go.tuiter/src/domain"
)

type TweetManager struct {
	TweetsMap    map[string][]*domain.Tweet
	FollowersMap map[string][]string
	NextId       int
}

func NewTweetManager() TweetManager {
	tweetManager := TweetManager{TweetsMap: make(map[string][]*domain.Tweet), FollowersMap: make(map[string][]string), NextId: 0}
	return tweetManager
}

func (tm *TweetManager) ResetService() {
	tm.TweetsMap = make(map[string][]*domain.Tweet)
	tm.FollowersMap = make(map[string][]string)
	tm.NextId = 0
}

func (tm *TweetManager) PublishTweet(newTweet *domain.Tweet) (int, error) {
	if newTweet.User == "" {
		return -1, fmt.Errorf("user is required")
	}
	if newTweet.Text == "" {
		return -1, fmt.Errorf("text is required")
	} else if len(newTweet.Text) > 140 {
		return -1, fmt.Errorf("text exceeds 140 characters")
	}

	tweets, ok := tm.TweetsMap[newTweet.User]
	if !ok {
		tm.TweetsMap[newTweet.User] = make([]*domain.Tweet, 0)
		tweets = tm.TweetsMap[newTweet.User]
	}
	tweets = append(tweets, newTweet)
	tm.TweetsMap[newTweet.User] = tweets
	newTweet.Id = tm.NextId
	tm.NextId++
	return newTweet.Id, nil
}

func (tm *TweetManager) GetTweetById(id int) *domain.Tweet {
	for _, tweets := range tm.TweetsMap {
		for _, tweet := range tweets {
			if tweet.Id == id {
				return tweet
			}
		}
	}
	return nil
}

func (tm *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	tweets, ok := tm.TweetsMap[user]
	if ok {
		return tweets
	}
	return nil
}

func (tm *TweetManager) GetTweets() []*domain.Tweet {
	allTweets := make([]*domain.Tweet, 0)
	for _, tweets := range tm.TweetsMap {
		allTweets = append(allTweets, tweets...)
	}
	return allTweets
}

func (tm *TweetManager) GetLatestTweet() *domain.Tweet {
	return tm.GetTweetById(tm.NextId - 1)
}

func (tm *TweetManager) CountTweetsByUser(user string) int {
	tweets, ok := tm.TweetsMap[user]
	if ok {
		return len(tweets)
	}
	return -1
}

func (tm *TweetManager) TweetCount() int {
	return tm.NextId
}

func (tm *TweetManager) DeleteTweets() {
	tm.TweetsMap = make(map[string][]*domain.Tweet)
	tm.NextId = 0
}

func (tm *TweetManager) Follow(followingUser, newFollowedUser string) error {
	if tm.CountTweetsByUser(newFollowedUser) == -1 {
		return fmt.Errorf("both users must exist")
	}

	followed, ok := tm.FollowersMap[followingUser]

	if !ok {
		tm.FollowersMap[followingUser] = make([]string, 0)
		followed = tm.FollowersMap[followingUser]
	} else {
		if contains(followed, newFollowedUser) {
			return fmt.Errorf("%s is already being followed by %s", followingUser, newFollowedUser)
		}

	}

	tm.FollowersMap[followingUser] = append(followed, newFollowedUser)

	return nil
}

func (tm *TweetManager) GetTimeline(user string) []*domain.Tweet {
	timeline := make([]*domain.Tweet, 0)
	followedUsers, ok := tm.FollowersMap[user]

	if !ok {
		return nil
	}

	for _, followedUser := range followedUsers {
		timeline = append(timeline, tm.TweetsMap[followedUser]...)
	}
	sort.Slice(timeline, func(i, j int) bool { return timeline[i].Date.Before(*timeline[j].Date) })
	return timeline
}

func contains(stringSlice []string, toFind string) bool {
	for _, s := range stringSlice {
		if s == toFind {
			return true
		}
	}
	return false
}

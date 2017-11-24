package service

import (
	"fmt"
	"sort"

	"github.com/alangberg/go.tuiter/src/domain"
)

type TweetManager struct {
	TweetsMap       map[string][]*domain.Tweet
	FollowersMap    map[string][]*domain.User
	RegisteredUsers []*domain.User
	NextId          int
}

func contains(usersSlice []*domain.User, toFind *domain.User) bool {
	for _, u := range usersSlice {
		if u.Username == toFind.Username {
			return true
		}
	}
	return false
}

func (tm *TweetManager) isRegistered(user *domain.User) bool {
	return contains(tm.RegisteredUsers, user)
}

func (tm *TweetManager) checkValidUser(newTweet *domain.Tweet) error {
	user := newTweet.User
	if user == nil {
		return fmt.Errorf(domain.EmptyUserErrorMessage)
	}

	if !tm.isRegistered(user) {
		return fmt.Errorf(domain.UnregisteredUserErrorMessage)
	}

	return nil
}

func (tm *TweetManager) checkTweetTextIsNotEmpty(newTweet *domain.Tweet) error {
	text := newTweet.Text
	if text == "" {
		return fmt.Errorf(domain.EmptyTextErrorMessage)
	}

	return nil
}

func (tm *TweetManager) checkValidTweetLenght(newTweet *domain.Tweet) error {
	text := newTweet.Text
	if len(text) > 140 {
		return fmt.Errorf(domain.ExceededLenghtErrorMessage)
	}

	return nil

}

func NewTweetManager() *TweetManager {
	tweetManager := TweetManager{TweetsMap: make(map[string][]*domain.Tweet), FollowersMap: make(map[string][]*domain.User), NextId: 0}
	return &tweetManager
}

func (tm *TweetManager) ResetService() {
	tm.TweetsMap = make(map[string][]*domain.Tweet)
	tm.FollowersMap = make(map[string][]*domain.User)
	tm.NextId = 0
}

func (tm *TweetManager) RegisterUser(newUser *domain.User) {
	if !tm.isRegistered(newUser) {
		tm.RegisteredUsers = append(tm.RegisteredUsers, newUser)
	}
}

func (tm *TweetManager) PublishTweet(newTweet *domain.Tweet) (int, error) {

	errValidUser := tm.checkValidUser(newTweet)
	if errValidUser != nil {
		return -1, errValidUser
	}

	errEmptyText := tm.checkTweetTextIsNotEmpty(newTweet)
	if errEmptyText != nil {
		return -1, errEmptyText
	}

	errTweetTextLenght := tm.checkValidTweetLenght(newTweet)
	if errTweetTextLenght != nil {
		return -1, errTweetTextLenght
	}

	tweets, ok := tm.TweetsMap[newTweet.User.Username]
	if !ok {
		tm.TweetsMap[newTweet.User.Username] = make([]*domain.Tweet, 0)
		tweets = tm.TweetsMap[newTweet.User.Username]
	}
	tweets = append(tweets, newTweet)
	tm.TweetsMap[newTweet.User.Username] = tweets
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

func (tm *TweetManager) GetTweetsByUser(user *domain.User) []*domain.Tweet {
	tweets, ok := tm.TweetsMap[user.Username]
	if ok {
		sort.Slice(tweets, func(i, j int) bool { return tweets[i].Date.Before(*tweets[j].Date) })
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

func (tm *TweetManager) CountTweetsByUser(user *domain.User) int {
	tweets, ok := tm.TweetsMap[user.Username]
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

func (tm *TweetManager) Follow(followingUser, newFollowedUser *domain.User) error {
	followerIsRegistered := tm.isRegistered(followingUser)
	followedIsRegistered := tm.isRegistered(newFollowedUser)

	if !(followedIsRegistered && followerIsRegistered) {
		return fmt.Errorf(domain.UnregisteredUserErrorMessage)
	}

	followed, ok := tm.FollowersMap[followingUser.Username]

	if !ok {
		tm.FollowersMap[followingUser.Username] = make([]*domain.User, 0)
		followed = tm.FollowersMap[followingUser.Username]
	} else {
		if contains(followed, newFollowedUser) {
			return fmt.Errorf(domain.AlreadyFollowingErrorMessage)
		}

	}

	tm.FollowersMap[followingUser.Username] = append(followed, newFollowedUser)

	return nil
}

func (tm *TweetManager) GetTimeline(user *domain.User) []*domain.Tweet {
	timeline := make([]*domain.Tweet, 0)
	followedUsers, ok := tm.FollowersMap[user.Username]

	if !ok {
		return nil
	}

	for _, followedUser := range followedUsers {
		timeline = append(timeline, tm.TweetsMap[followedUser.Username]...)
	}
	sort.Slice(timeline, func(i, j int) bool { return timeline[i].Date.Before(*timeline[j].Date) })
	return timeline
}

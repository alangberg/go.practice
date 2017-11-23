package service_test

import (
	"testing"

	"github.com/alangberg/go.tuiter/src/domain"
	"github.com/alangberg/go.tuiter/src/service"
)

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.User != user && tweet.Text != text && tweet.Id != id {
		t.Errorf("Expected tweet: \n Id: %d \n, User: %s \n, Text: %s, \n but is \n Id: %d \n, User: %s \n, Text: %s \n",
			id, user, text, tweet.Id, tweet.User, tweet.Text)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestPublishedTweetIsSaved(t *testing.T) {
	service.InitializeService()

	user := "grupoEsfera"
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)

	var tweetId int
	tweetId, _ = service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()

	isValidTweet(t, publishedTweet, tweetId, user, text)
}

func TestCleanTweetDeletesTweet(t *testing.T) {
	service.InitializeService()

	user := "grupoEsfera"
	text := "Tweet to be deleted"

	tweet := domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	service.DeleteTweets()

	if service.GetTweet() != nil {
		t.Error("No tweets expected")
	}

}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	service.InitializeService()

	var user string
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)

	var err error
	_, err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}

	if service.TweetCount() != 0 {
		t.Error("Expected number of tweets is zero")
	}

}

func TestCanNotPublishTweetWithoutText(t *testing.T) {
	service.InitializeService()

	var text string
	user := "grupoEsfera"
	tweet := domain.NewTweet(user, text)

	_, err := service.PublishTweet(tweet)

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}

	if service.TweetCount() != 0 {
		t.Error("Expected number of tweets is zero")
	}

}

func TestCanNotPublishTweetsLongerThan140Characters(t *testing.T) {

	service.InitializeService()

	text := `Go became a public open source project on November 10, 2009. After a couple of years of very active design and development, stability was called for and Go 1 was released on March 28, 2012. Go 1, which includes a language specification, standard libraries, and custom tools, provides a stable foundation for creating reliable products, projects, and publications.
	With that stability established, we are using Go to develop programs, products, and tools rather than actively changing the language and libraries. In fact, the purpose of Go 1 is to provide long-term stability. Backwards-incompatible changes will not be made to any Go 1 point release. We want to use what we have to learn how a future version of Go might look, rather than to play with the language underfoot. `
	user := "grupoEsfera"

	tweet := domain.NewTweet(user, text)

	_, err := service.PublishTweet(tweet)

	if err == nil || err.Error() != "text can not be longer than 140 characters" {
		t.Error("Expected error is text can not be longer than 140 characters")
	}

	if service.TweetCount() != 0 {
		t.Error("Expected number of tweets is zero")
	}

}
func TestCanRetrieveTweetById(t *testing.T) {
	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	count := service.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	firstId, _ := service.PublishTweet(tweet)
	secondId, _ := service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	tweets := service.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func TestFollowUser(t *testing.T) {
	service.InitializeService()

	followedUser1 := "FollowMe1"
	followedUser2 := "FollowMe2"
	notFollowedUser := "Nobody loves me"
	followingUser := "Following"

	text1 := "hi"
	tweet1 := domain.NewTweet(followedUser1, text1)

	text2 := "bye"
	tweet2 := domain.NewTweet(followedUser2, text2)

	text3 := "bu-hu"
	tweet3 := domain.NewTweet(notFollowedUser, text3)

	tweet1Id, _ := service.PublishTweet(tweet1)
	tweet2Id, _ := service.PublishTweet(tweet2)
	service.PublishTweet(tweet3)

	service.Follow(followingUser, followedUser1)
	service.Follow(followingUser, followedUser2)

	timeline := service.GetTimeline(followingUser)

	if len(timeline) != 2 {
		t.Errorf("Expected 2 tweets but was %d", len(timeline))
		return
	}

	if !isValidTweet(t, timeline[0], tweet1Id, followedUser1, text1) {
		t.Errorf("Expected first tweet to be from %s but was from %s", followedUser1, timeline[0].User)
		return
	}

	if !isValidTweet(t, timeline[1], tweet2Id, followedUser2, text2) {
		t.Errorf("Expected first tweet to be from %s but was from %s", followedUser2, timeline[1].User)
		return
	}

}

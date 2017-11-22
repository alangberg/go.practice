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

	user := "grupoEsfera"
	text := "This is my first tweet"
	tweet := domain.NewTweet(user, text)

	var tweetId int
	tweetId, _ = service.PublishTweet(tweet)

	publishedTweets := service.GetTweets()
	publishedTweet := publishedTweets[0]

	isValidTweet(t, publishedTweet, tweetId, user, text)
}

func TestCleanTweetDeletesTweet(t *testing.T) {
	user := "grupoEsfera"
	text := "Tweet to be deleted"

	tweet := domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	service.DeleteTweets()

	if len(service.GetTweets()) != 0 {
		t.Error("No tweets expected")
	}

}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
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

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	var firstTweetId, secondTweetId int
	// Operation
	firstTweetId, _ = service.PublishTweet(tweet)
	secondTweetId, _ = service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstTweetId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondTweetId, user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {
	// Initialization

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

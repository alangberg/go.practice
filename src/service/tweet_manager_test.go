package service_test

import (
	"testing"

	"github.com/alangberg/go.tuiter/src/domain"
	"github.com/alangberg/go.tuiter/src/service"
)

var defUser, defSecondUser, defUnregisteredUser *domain.User
var defTweetText string
var defTweet *domain.Tweet
var tweetManager *service.TweetManager

func defaultUser() *domain.User {
	return domain.NewUser("defaultUser")
}

func defaultTweetText() string {
	return "Default tweet text"
}

func defaultTweet() *domain.Tweet {
	return &domain.Tweet{User: defaultUser(), Text: defaultTweetText()}
}

func TestMain(m *testing.M) {
	defUser = defaultUser()
	defTweetText = defaultTweetText()
	defTweet = defaultTweet()
	defSecondUser = &domain.User{Username: "defSecondUser"}
	defUnregisteredUser = &domain.User{Username: "defUnregisteredUser"}

	m.Run()
}

func TestUnregisteredUserCanNotPublishTweet(t *testing.T) {
	tweetManager = service.NewTweetManager()

	// Operation
	_, err := tweetManager.PublishTweet(defTweet)

	// Validation

	if tweetManager.TweetCount() != 0 {
		t.Errorf("Did not expect tweet to be pubished")
	}

	if err == nil || err.Error() != domain.UnregisteredUserErrorMessage {
		errorMessage := "Expected error '" + domain.UnregisteredUserErrorMessage + "'"
		if err != nil {
			errorMessage = errorMessage + "but was '" + err.Error() + "'"
		}
		t.Errorf(errorMessage)
	}

}

func TestUnregisteredUserCanNotFollow(t *testing.T) {
	tweetManager = service.NewTweetManager()

	err := tweetManager.Follow(defUnregisteredUser, defUser)

	// Validation
	if err != nil || err.Error() != domain.UnregisteredUserErrorMessage {
		t.Errorf("Expected error '%s' but was '%s'", domain.UnregisteredUserErrorMessage, err.Error())
	}
}

func TestPublishedTweetIsSaved(t *testing.T) {
	tweetManager = service.NewTweetManager()

	// Operation
	id, _ := tweetManager.PublishTweet(defTweet)

	// Validation
	publishedTweet := tweetManager.GetLatestTweet()

	isValidTweet(t, publishedTweet, id, defUser, defTweetText)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	tweetManager = service.NewTweetManager()

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(defTweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	tweetManager = service.NewTweetManager()

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(defTweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	tweetManager = service.NewTweetManager()

	longText := `The Go project has grown considerably with over half a million users and community members
	   all over the world. To date all community oriented activities have been organized by the community
	   with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet := domain.NewTweet(defUser, longText)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}

	if tweetManager.TweetCount() != 0 {
		t.Error("Did not expect tweet to be pubished")
	}
}
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	tweetManager = service.NewTweetManager()

	secondTweetText := "This is my second tweet"

	firstTweet := defTweet
	secondTweet := domain.NewTweet(defUser, secondTweetText)

	// Operation
	firstId, _ := tweetManager.PublishTweet(firstTweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, defUser, defTweetText) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, defUser, secondTweetText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {
	tweetManager = service.NewTweetManager()

	// Operation
	id, _ := tweetManager.PublishTweet(defTweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, defUser, defTweetText)
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	tweetManager = service.NewTweetManager()

	// Initialization

	secondTweetText := "This is my second tweet"

	firstTweet := defTweet
	secondTweet := domain.NewTweet(defUser, secondTweetText)
	thirdTweet := domain.NewTweet(defSecondUser, secondTweetText)

	firstId, _ := tweetManager.PublishTweet(firstTweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(defUser)

	// Validation

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, defUser, defTweetText) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, defUser, secondTweetText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user *domain.User, text string) bool {

	if tweet.Id != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.Id)
	}

	if tweet.User != user && tweet.Text != text {
		t.Errorf("Expected tweet from %s: %s \nbut is from %s: %s",
			user.Username, text, tweet.User.Username, tweet.Text)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

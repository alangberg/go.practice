package service_test

import (
	"testing"

	"github.com/alangberg/go.tuiter/src/domain"
	"github.com/alangberg/go.tuiter/src/service"
)

var defUser, defSecondUser, defUnregisteredUser *domain.User
var defTweetText string
var defTweet domain.Tweet
var tweetManager *service.TweetManager

func defaultUser() *domain.User {
	return domain.NewUser("defaultUser")
}

func defaultTweetText() string {
	return "Default tweet text"
}

func defaultTextTweet() domain.Tweet {
	return domain.NewTextTweet(defaultUser(), defaultTweetText())
}

func TestMain(m *testing.M) {
	defUser = defaultUser()
	defTweetText = defaultTweetText()
	defTweet = defaultTextTweet()
	defSecondUser = &domain.User{Username: "defSecondUser"}
	defUnregisteredUser = &domain.User{Username: "defUnregisteredUser"}

	m.Run()
}

func TestUnregisteredUserCanNotPublishTweet(t *testing.T) {
	tweetManager = service.NewTweetManager("memory")

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
	tweetManager = service.NewTweetManager("memory")

	err := tweetManager.Follow(defUnregisteredUser, defUser)

	// Validation
	if err == nil || err.Error() != domain.UnregisteredUserErrorMessage {
		errorMessage := "Expected error '" + domain.UnregisteredUserErrorMessage + "'"
		if err != nil {
			errorMessage = errorMessage + "but was '" + err.Error() + "'"
		}
		t.Errorf(errorMessage)
	}
}

func TestPublishedTweetIsSaved(t *testing.T) {
	tweetManager = service.NewTweetManager("memory")
	tweetManager.RegisterUser(defUser)
	// Operation
	id, _ := tweetManager.PublishTweet(defTweet)

	// Validation
	publishedTweet := tweetManager.GetLatestTweet()

	isValidTweet(t, publishedTweet, id, defUser, defTweetText)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	tweetManager = service.NewTweetManager("memory")
	tweetManager.RegisterUser(defUser)
	// Operation
	var err error
	_, err = tweetManager.PublishTweet(defTweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	tweetManager = service.NewTweetManager("memory")
	tweetManager.RegisterUser(defUser)

	emptyTweet := domain.NewTextTweet(defUser, "")

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(emptyTweet)

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
	tweetManager = service.NewTweetManager("memory")
	tweetManager.RegisterUser(defUser)
	longText := `The Go project has grown considerably with over half a million users and community members
	   all over the world. To date all community oriented activities have been organized by the community
	   with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet := domain.NewTextTweet(defUser, longText)

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
	tweetManager = service.NewTweetManager("memory")
	tweetManager.RegisterUser(defUser)
	secondTweetText := "This is my second tweet"

	firstTweet := defTweet
	secondTweet := domain.NewTextTweet(defUser, secondTweetText)

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
	tweetManager = service.NewTweetManager("memory")
	tweetManager.RegisterUser(defUser)
	// Operation
	id, _ := tweetManager.PublishTweet(defTweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, defUser, defTweetText)
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	tweetManager = service.NewTweetManager("memory")
	tweetManager.RegisterUser(defUser)
	tweetManager.RegisterUser(defSecondUser)
	// Initialization

	secondTweetText := "This is my second tweet"

	firstTweet := defTweet
	secondTweet := domain.NewTextTweet(defUser, secondTweetText)
	thirdTweet := domain.NewTextTweet(defSecondUser, secondTweetText)

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

func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user *domain.User, text string) bool {

	if tweet.GetId() != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.GetId())
	}

	if tweet.GetUser() != user && tweet.GetText() != text {
		t.Errorf("Expected tweet from %s: %s \nbut is from %s: %s",
			user.Username, text, tweet.GetUser().Username, tweet.GetText())
		return false
	}

	if tweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestPluginsHaveNoErrors(t *testing.T) {
	tweetManager = service.NewTweetManager("memory")
	fbPlugin := &service.FacebookPlugin{}

	tweetManager.AddPlugin(fbPlugin)

	// Operation
	_, err := tweetManager.PublishTweet(defTweet)

	// Validation
	if err != nil && err.Error() != "Plugin error" {
		t.Error("Expected error is user is required")
	}
}

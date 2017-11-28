package service_test

import (
	"testing"

	"github.com/alangberg/go.tuiter/src/domain"
	"github.com/alangberg/go.tuiter/src/service"
)

func TestCanWriteATweet(t *testing.T) {

	// Initialization
	tweet := defTweet
	tweet2 := domain.NewTextTweet(defSecondUser, defTweetText)

	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetsToWrite := make(chan domain.Tweet)
	quit := make(chan bool)

	go tweetWriter.WriteTweet(tweetsToWrite, quit)

	// Operation
	tweetsToWrite <- tweet
	tweetsToWrite <- tweet2
	close(tweetsToWrite)

	<-quit

	// Validation
	if memoryTweetWriter.GetTweets()[0] != tweet {
		t.Errorf("A tweet in the writer was expected")
	}

	if memoryTweetWriter.GetTweets()[1] != tweet2 {
		t.Errorf("A tweet in the writer was expected")
	}
}

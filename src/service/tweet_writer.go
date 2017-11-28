package service

import (
	"os"

	"github.com/alangberg/go.tuiter/src/domain"
)

type MemoryTweetWriter struct {
	tweets []domain.Tweet
}

type FileTweetWriter struct {
	file *os.File
}

type ChannelTweetWriter struct {
	tweetWriter domain.TweetWriter
}

//MemoryTweetWriter

func (mtw *MemoryTweetWriter) WriteTweet(newTweet domain.Tweet) {
	mtw.tweets = append(mtw.tweets, newTweet)
}

func (mtw *MemoryTweetWriter) GetTweets() []domain.Tweet {
	return mtw.tweets
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	writer := new(MemoryTweetWriter)
	writer.tweets = make([]domain.Tweet, 0)
	return writer
}

//FileTweetWriter

func (ftw *FileTweetWriter) WriteTweet(newTweet domain.Tweet) {
	if ftw.file != nil {
		byteSlice := []byte(newTweet.PrintableTweet() + "\n")
		ftw.file.Write(byteSlice)
	}
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)
	writer.file = file

	return writer
}

//ChannelTweetWriter

func (ctw *ChannelTweetWriter) WriteTweet(tweetsToWrite chan domain.Tweet, quit chan bool) {
	tweet, open := <-tweetsToWrite

	for open {
		ctw.tweetWriter.WriteTweet(tweet)
		tweet, open = <-tweetsToWrite
	}
	quit <- true
}

func NewChannelTweetWriter(tweetWriter domain.TweetWriter) *ChannelTweetWriter {
	channelWriter := new(ChannelTweetWriter)
	channelWriter.tweetWriter = tweetWriter
	return channelWriter
}

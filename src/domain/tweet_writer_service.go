package domain

type TweetWriter interface {
	WriteTweet(newTweet Tweet)
}

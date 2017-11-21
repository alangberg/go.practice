package service

var tweet string

func PublishTweet(newTweet string) {
	tweet = newTweet
}

func GetTweet() string {
	return tweet
}

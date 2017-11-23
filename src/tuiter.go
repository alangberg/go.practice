package main

import (
	"github.com/abiosoft/ishell"
	"github.com/alangberg/go.tuiter/src/domain"
	"github.com/alangberg/go.tuiter/src/service"
)

func printTweets(tweets []*domain.Tweet, c *ishell.Context) {
	if len(tweets) == 0 {
		c.Println("No tweets to show :(")
	} else {
		for i := 0; i < len(tweets); i++ {
			c.Println("Tweet ID:", tweets[i].Id)
			c.Println("User:", tweets[i].User)
			c.Println("Content:", tweets[i].Text)
			c.Println("Date:", tweets[i].Date.Format("02-01-2006 15:04:05"))
			c.Println()
		}
	}
}

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tuit3r >> ")
	shell.Print("Type 'help' to know commands\n")

	tm := service.NewTweetManager()

	/*	shell.Print("Hello! Please enter your username:\n")
		defer shell.ShowPrompt(true)
		user := shell.ReadLine()
	*/
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Please enter your username:")
			user := c.ReadLine()

			c.Print("Write your tweet:")

			text := c.ReadLine()
			tweet := domain.NewTweet(user, text)

			_, err := tm.PublishTweet(tweet)

			if err != nil {
				c.Println(err.Error())
			} else {
				c.Println("Tweet sent.")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsFromUser",
		Help: "Shows all tweets from a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Please enter a username:")
			user := c.ReadLine()

			tweets := tm.GetTweetsByUser(user)

			printTweets(tweets, c)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showallTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Please enter a username:")
			user := c.ReadLine()

			tweets := tm.GetTweetsByUser(user)

			printTweets(tweets, c)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "deleteTweets",
		Help: "Deletes all existing tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tm.DeleteTweets()

			c.Println("Tweets Deleted.")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countUserTweets",
		Help: "Count all existing tweets from a given user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Please enter a username:")
			user := c.ReadLine()

			tweetCount := tm.CountTweetsByUser(user)

			c.Printf("The user %s has %d tweets. \n", user, tweetCount)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "follow",
		Help: "us3r A follou user B",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Please enter a follower:")
			follower := c.ReadLine()

			c.Print("Please enter a user to follow:")
			followed := c.ReadLine()

			tm.Follow(follower, followed)

			c.Printf("%s is now following %s! :D \n", follower, followed)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTimeline",
		Help: "Count all existing tweets from a given user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Please enter a username:")
			user := c.ReadLine()

			tweets := tm.GetTimeline(user)
			c.Printf("The user %s has %d tweets in his\\her timeline. \n", user, len(tweets))

			printTweets(tweets, c)

			return
		},
	})

	shell.Run()

}

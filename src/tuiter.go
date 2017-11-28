package main

import (
	"github.com/abiosoft/ishell"
	"github.com/alangberg/go.tuiter/src/domain"
	"github.com/alangberg/go.tuiter/src/rest"
	"github.com/alangberg/go.tuiter/src/service"
)

func printTweets(tweets []domain.Tweet, c *ishell.Context) {
	if len(tweets) == 0 {
		c.Println("No tweets to show :(")
	} else {
		for i := 0; i < len(tweets); i++ {
			c.Println("Tweet ID:", tweets[i].GetId())
			c.Println("User:", tweets[i].GetUser().Username)
			c.Println("Content:", tweets[i].GetText())
			c.Println("Date:", tweets[i].GetDate().Format("02-01-2006 15:04:05"))
			c.Println()
		}
	}
}

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tuit3r >> ")
	shell.Print("Type 'help' to know commands\n")

	tm := service.NewTweetManager("file")

	ginServer := rest.NewGinServer(tm)
	ginServer.StartGinServer()

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
			username := c.ReadLine()
			user := domain.NewUser(username)
			c.Print("Write your tweet:")

			text := c.ReadLine()
			tweet := domain.NewTextTweet(user, text)

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
			username := c.ReadLine()
			user := domain.NewUser(username)

			tweets := tm.GetTweetsByUser(user)

			printTweets(tweets, c)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tm.GetTweets()

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

			username := c.ReadLine()
			user := domain.NewUser(username)
			tweetCount := tm.CountTweetsByUser(user)

			c.Printf("The user %s has %d tweets. \n", user.Username, tweetCount)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "follow",
		Help: "us3r A follou user B",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Please enter a follower:")
			followerUsername := c.ReadLine()
			follower := domain.NewUser(followerUsername)
			c.Print("Please enter a user to follow:")
			followedUsername := c.ReadLine()
			followed := domain.NewUser(followedUsername)

			tm.Follow(follower, followed)

			c.Printf("%s is now following %s! :D \n", follower.Username, followed.Username)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTimeline",
		Help: "Count all existing tweets from a given user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Please enter a username:")
			username := c.ReadLine()
			user := domain.NewUser(username)

			tweets := tm.GetTimeline(user)
			c.Printf("The user %s has %d tweets in his\\her timeline. \n", user.Username, len(tweets))

			printTweets(tweets, c)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "registerUser",
		Help: "Register a new us3r",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Username:")
			newUsername := c.ReadLine()
			newUser := domain.NewUser(newUsername)

			tm.RegisterUser(newUser)

			c.Printf("Welcome %s! \n", newUsername)

			return
		},
	})

	shell.Run()

}

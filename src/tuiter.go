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
			c.Println("Tweet ID:", i)
			c.Println("User:", tweets[i].User)
			c.Println("Content:", tweets[i].Text)
			c.Println("Date:", tweets[i].Date.Format("02-01-2006 15:04:05"))
			c.Println()
		}
	}
}

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tuiter >> ")
	shell.Print("Type 'help' to know commands\n")

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

			err := service.PublishTweet(tweet)

			if err != nil {
				c.Println(err.Error())
			} else {
				c.Println("Tweet sent.")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			printTweets(tweets, c)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "deleteTweets",
		Help: "Deletes all existing tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			service.DeleteTweets()

			c.Println("Tweets Deleted.")

			return
		},
	})

	shell.Run()

}

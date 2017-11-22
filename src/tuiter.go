package main

import (
	"github.com/abiosoft/ishell"
	"github.com/alangberg/go.tuiter/src/domain"
	"github.com/alangberg/go.tuiter/src/service"
)

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

			service.PublishTweet(tweet)

			c.Print("Tweet sent. \n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "deleteTweet",
		Help: "Deletes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			service.DeleteTweet()

			c.Println("Tweet Deleted.\n")

			return
		},
	})

	shell.Run()

}

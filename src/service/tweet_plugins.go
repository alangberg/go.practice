package service

import (
	"fmt"
)

type FacebookPlugin struct{}
type GooglePlusPlugin struct{}

func (fp *FacebookPlugin) RunPlugin() error {
	fmt.Println("Your Tweet has been shared on Facebook.")
	return nil
}

func (fp *FacebookPlugin) GetPluginName() string {
	return "FacebookTweetPlugin"
}

func (ip *GooglePlusPlugin) RunPlugin() error {
	fmt.Println("Your Tweet has been shared on Google+.")
	return nil
}

func (ip *GooglePlusPlugin) GetPluginName() string {
	return "Google+TweetPlugin"
}

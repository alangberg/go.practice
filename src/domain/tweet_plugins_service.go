package domain

type TweetPlugin interface {
	RunPlugin() error
	GetPluginName() string
}

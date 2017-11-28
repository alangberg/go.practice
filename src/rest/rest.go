package rest

import (
	"net/http"

	"github.com/alangberg/go.tuiter/src/domain"
	"github.com/alangberg/go.tuiter/src/service"
	"github.com/gin-gonic/gin"
)

type GinTweet struct {
	Username string
	Text     string
}

type GinServer struct {
	tweetManager *service.TweetManager
}

func NewGinServer(tweetManager *service.TweetManager) *GinServer {
	return &GinServer{tweetManager}
}

func (server *GinServer) StartGinServer() {

	router := gin.Default()

	router.GET("/listTweets", server.listTweets)
	router.GET("/listTweets/:username", server.listTweets)
	router.POST("publishTweet", server.publishTweet)

	go router.Run()
}

func (server *GinServer) listTweets(c *gin.Context) {

	c.JSON(http.StatusOK, server.tweetManager.GetTweets())
}

func (server *GinServer) getTweetsByUser(c *gin.Context) {

	username := c.Param("username")
	user := domain.NewUser(username)

	c.JSON(http.StatusOK, server.tweetManager.GetTweetsByUser(user))
}

func (server *GinServer) publishTweet(c *gin.Context) {

	var tweetdata GinTweet
	c.Bind(&tweetdata)

	user := domain.NewUser(tweetdata.Username)

	tweetToPublish := domain.NewTextTweet(user, tweetdata.Text)

	id, err := server.tweetManager.PublishTweet(tweetToPublish)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ Id int }{id})
	}
}

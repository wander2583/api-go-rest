package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wander2583/api-go-rest/api/entities"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}

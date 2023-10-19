package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wander2583/api-go-rest/api/controllers"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {

	tweetController := controllers.NewTweetController()

	v1 := router.Group("/v1")
	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweets", tweetController.Create)
	}

	return v1
}

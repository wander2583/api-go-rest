package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wander2583/api-go-rest/api/routes"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)
	app.Run("localhost:3001")
}

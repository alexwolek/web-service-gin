package main

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controllers.AddAlbumsController(router)

	router.Run("localhost:8080")
}

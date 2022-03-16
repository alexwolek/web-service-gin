package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("AUDIENCE", "api")
	os.Setenv("AUTHORITY", "https://demo.identityserver.io")
	os.Setenv("PORT", "8080")

	router := gin.Default()
	router.Use(middleware.EnsureValidToken())

	controllers.AddAlbumsController(router)

	router.Run()
}

package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	if os.Getenv("AUDIENCE") == "" {
		os.Setenv("AUDIENCE", "api")
	}
	if os.Getenv("AUTHORITY") == "" {
		os.Setenv("AUTHORITY", "https://demo.identityserver.io")
	}
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}

	router := gin.Default()
	router.Use(middleware.EnsureValidToken())

	controllers.AddAlbumsController(router)

	router.Run()
}

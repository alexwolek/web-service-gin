package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHealthCheck(engine *gin.Engine, handlers ...gin.HandlerFunc) {
	health := engine.Group("/health", handlers...)
	{
		health.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "Healthy")
		})
	}
}

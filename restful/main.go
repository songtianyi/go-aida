package main

import (
	"github.com/gin-gonic/gin"
	"github.com/songtianyi/go-aida/restful/logic"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(gin.HandlerFunc(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}))

	router.GET("/create", logic.Create)
	router.GET("/status", logic.Status)
	router.PUT("/enable", logic.Enable)
	router.PUT("/disable", logic.Disable)

	router.Run()
}

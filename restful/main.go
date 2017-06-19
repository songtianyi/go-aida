package main

import (
	"github.com/gin-gonic/gin"
	"github.com/songtianyi/go-aida/restful/logic"
)

func main() {
	router := gin.Default()

	router.Use(gin.HandlerFunc(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}))

	router.GET("/create", logic.Create)
	router.GET("/status", logic.Status)
	router.PATCH("/enable", logic.Enable)
	router.PATCH("/disable", logic.Disable)

	router.Run()
}

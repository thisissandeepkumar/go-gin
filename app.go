package main

import (
	"fmt"
	"net/http"

	"github.com/gn-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/ping", func(c *gin.Context) {
		// message := c.PostForm("message")
		fmt.Println(c.Request.Body)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

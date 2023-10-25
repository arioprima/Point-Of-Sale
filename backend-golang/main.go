package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/push", func(c *gin.Context) {
		c.String(200, "push")
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

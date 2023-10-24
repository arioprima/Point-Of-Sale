package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, World!")
	})

	router.GET("/satu", func(ctx *gin.Context) {
		ctx.String(200, "Hello, satu!")
	})
	router.GET("/dua", func(ctx *gin.Context) {
		ctx.String(200, "Hello, satu!")
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

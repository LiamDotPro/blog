package main

import (
	"fmt"
	"os"

	data "github.com/choskyo/blog/data"
	"github.com/choskyo/blog/posts"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "こんにちは",
		})
	})

	data.Start()

	posts.Setup(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Fprintf(os.Stderr, "Error running server: %v", err)
	}
}

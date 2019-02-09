package main

import (
	"fmt"
	"os"

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

	p := posts.Post{
		Title: "Test",
	}

	fmt.Printf(p.Title)

	if err := router.Run(":8080"); err != nil {
		fmt.Fprintf(os.Stderr, "Error running server: %v", err)
	}
}

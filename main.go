package main

import (
	"fmt"
	"os"

	"github.com/choskyo/blog/auth"
	"github.com/choskyo/blog/data"
	"github.com/choskyo/blog/posts"
	"github.com/choskyo/blog/users"
	"github.com/gin-gonic/gin"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(cors())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "こんにちは",
		})
	})

	data.Start()

	// Initialise routes + migrate db
	posts.Setup(router)
	users.Setup(router)
	auth.Setup(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Fprintf(os.Stderr, "Error running server: %v", err)
	}
}

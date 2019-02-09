package posts

import (
	"time"

	data "github.com/choskyo/blog/data"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Post ...
type Post struct {
	gorm.Model

	Title     string
	Body      string
	Timestamp time.Time
}

// Setup initialises post routes + db table
func Setup(router *gin.Engine) {
	router.GET("/api/posts", get)
}

func get(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": data.DataString,
	})
}

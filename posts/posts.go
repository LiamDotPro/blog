package posts

import (
	"strconv"
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
	router.GET("/api/posts", getPosts)
	router.GET("/api/posts/:id", getPost)

	if !data.Connection.HasTable(Post{}) {
		data.Connection.CreateTable(Post{})
	}
}

func getPosts(c *gin.Context) {
	var posts []Post

	data.Connection.Find(&posts)

	if len(posts) == 0 {
		c.Status(204)
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func getPost(c *gin.Context) {
	param := c.Param("id")

	if len(param) == 0 {
		c.JSON(400, gin.H{
			"err": "url param 'id' required but not found",
		})
		return
	}

	postID, err := strconv.Atoi(param)

	if err != nil {
		c.JSON(400, gin.H{
			"err": "failed parsing postId",
		})
	}

	var post Post

	data.Connection.First(&post, postID)

	c.JSON(200, gin.H{
		"post": post,
	})
}

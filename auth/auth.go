package auth

import (
	"github.com/choskyo/blog/data"
	"github.com/choskyo/blog/users"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Setup initalises routes + tables
func Setup(r *gin.Engine) {
	r.POST("/api/auth/login", login)
}

func login(c *gin.Context) {
	var login loginModel

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"message": "Incorrect format"})
		return
	}

	var user users.User

	if data.Connection.First(&user, &users.User{Name: login.Name}).RecordNotFound() {
		c.JSON(400, gin.H{"message": "Incorrect username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(user.Password)); err != nil {
		c.JSON(400, gin.H{"message": "Incorrect username or password"})
		return
	}

	//TODO: finish login
	c.JSON(200, gin.H{
		"message": "Success",
	})
}

package auth

import (
	"net/http"

	"github.com/choskyo/blog/data"
	"github.com/choskyo/blog/users"
	"github.com/gin-gonic/gin"
	"github.com/wader/gormstore"
	"golang.org/x/crypto/bcrypt"
)

func isAuthorized(Store *gormstore.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

		sessionValues, err := Store.Get(c.Request, "connect.s.id")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to view this."})
			c.Abort()
		}

		if sessionValues.Values["Authorised"] != true {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to view this."})
			c.Abort()
		}

		// Pass the user id into the handler.
		c.Set("userId", sessionValues.Values["userId"])
	}
}

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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(400, gin.H{"message": "Incorrect username or password"})
		return
	}

	session, _ := data.Store.Get(c.Request, "session")
	session.Values["userID"] = user.ID
	session.Save(c.Request, c.Writer)

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

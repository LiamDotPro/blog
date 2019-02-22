package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/choskyo/blog/data"
	"github.com/choskyo/blog/users"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Setup initalises routes + tables
func Setup(r *gin.Engine) {
	r.POST("/api/auth/login", login)
	r.GET("/api/auth/test", testLogin)
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

	// Setup new session.
	session, err := data.Store.New(c.Request, "session")

	if err != nil {
		fmt.Println(err)
	}

	session.Values["Authorised"] = true
	session.Values["userId"] = user.ID

	data.Store.Save(c.Request, c.Writer, session)

	c.JSON(http.StatusOK, gin.H{
		"attempt": "uh ok",
		"message": "You have successfully logged into your account.",
	})
}

func testLogin(c *gin.Context) {
	session, err := data.Store.Get(c.Request, "session")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	val, ok := session.Values["userId"]

	fmt.Printf("%v\n\n", len(session.Values))

	if !ok {
		fmt.Fprintf(os.Stderr, "not ok!\n")
		c.Status(500)
		return
	}

	username, ok := val.(string)
	if !ok {
		fmt.Fprintf(os.Stderr, "casting was not ok!\n")
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"user": username,
	})
}

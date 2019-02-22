package users

import (
	"strconv"

	"github.com/choskyo/blog/data"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Setup initialises users routes + db table
func Setup(router *gin.Engine) {
	router.GET("/api/users", getUsers)
	router.GET("/api/users/:id", getUser)
	router.POST("/api/users", newUser)

	if !data.Connection.HasTable(User{}) {
		data.Connection.CreateTable(User{})

		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), 4)

		data.Connection.Create(&User{
			Name:     "admin",
			Password: string(hash),
			Admin:    true,
		})
	}

	data.Connection.AutoMigrate(&User{})
}

func getUsers(c *gin.Context) {
	var users []User

	data.Connection.Find(&users)

	if len(users) == 0 {
		c.Status(204)
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func getUser(c *gin.Context) {
	param := c.Param("id")

	if len(param) == 0 {
		c.JSON(400, gin.H{
			"err": "url param 'id' required but not found",
		})
		return
	}

	userID, err := strconv.Atoi(param)

	if err != nil {
		c.JSON(400, gin.H{
			"err": "failed parsing userId",
		})
	}

	var user User

	data.Connection.First(&user, userID)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func newUser(c *gin.Context) {
	var nu newUserModel

	if err := c.ShouldBind(&nu); err != nil {
		c.JSON(400, gin.H{
			"err": "Invalid format",
		})
		return
	}

	if validationErrors := validNewUser(nu); validationErrors != nil {
		c.JSON(400, gin.H{
			"errs": validationErrors,
		})
		return
	}

	if !data.Connection.First(&User{}, &User{Name: nu.Name}).RecordNotFound() {
		c.JSON(400, gin.H{
			"err": "Username taken",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), 4)
	if err != nil {
		c.JSON(500, gin.H{
			"err": "Couldn't hash password",
		})
		return
	}

	user := &User{
		Name:     nu.Name,
		Password: string(hash),
		Admin:    nu.Admin,
	}

	if err := data.Connection.Create(user).Error; err != nil {
		c.JSON(500, gin.H{
			"err": "Couldn't create new record in db",
		})
		return
	}

	c.JSON(200, gin.H{
		"id": user.ID,
	})
}

package posts

import (
	"time"

	users "github.com/choskyo/blog/users"
	"github.com/jinzhu/gorm"
)

// Post ...
type Post struct {
	gorm.Model

	Title     string
	Body      string
	Author    users.User
	AuthorID  int
	Timestamp time.Time
}

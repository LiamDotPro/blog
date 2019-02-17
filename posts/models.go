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
	AuthorID  users.User
	Timestamp time.Time
}

package posts

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Post ...
type Post struct {
	gorm.Model

	Title     string
	Body      string
	Timestamp time.Time
}

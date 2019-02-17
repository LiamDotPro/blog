package users

import "github.com/jinzhu/gorm"

// User ...
type User struct {
	gorm.Model

	Name     string
	Password string
	Admin    bool
}

type newUserModel struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Admin    bool   `json:"admin"`
}

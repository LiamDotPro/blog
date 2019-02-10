package data

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // req
)

// Connection global db connection
var Connection *gorm.DB

// Start opens db connection
func Start() {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=blog password=Test1234 sslmode=disable")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening DB failed: %v\n", err)
		os.Exit(1)
	}

	Connection = db
}

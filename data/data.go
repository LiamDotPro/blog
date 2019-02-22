package data

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // req
	"github.com/wader/gormstore"
)

// Connection global db connection
var Connection *gorm.DB

// Store user sessions
var Store *gormstore.Store

// Start opens db connection
func Start() {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=blog password=Test1234 sslmode=disable")

	store := gormstore.New(db, []byte(os.Getenv("SESSION_KEY")))
	store.SessionOpts.Secure = true

	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening DB failed: %v\n", err)
		os.Exit(1)
	}

	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)

	Connection = db
	Store = store
}

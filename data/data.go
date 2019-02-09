package data

import (
	"fmt"
	_ "os"

	_ "github.com/jinzhu/gorm"
)

// Connection global db connection
// var Connection *gorm.DB

// DataString is teststuff
var DataString string

// Start opens db connection
func Start() {
	// db, err := gorm.Open("postgres", "")

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Opening DB failed: %v\n", err)
	// }

	// Connection = db

	DataString = "ABCDEFG"
	fmt.Println("DB RUN")
	fmt.Println("DB " + DataString)
}

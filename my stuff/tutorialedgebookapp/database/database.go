package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

func InitDatabase() {
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("✔️  Connection Opened to Database")
}

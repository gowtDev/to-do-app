package data

import (
	"log"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	mu sync.Mutex
)

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&Todo{}, &User{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}
}

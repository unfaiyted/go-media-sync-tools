package db

import (
	"media-sync-tools/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("media-sync.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	err = db.AutoMigrate(&types.User{})

	if err != nil {
		panic("failed to migrate database")
	}

	return db
}

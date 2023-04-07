package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"api/domain"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&domain.Club{}); err != nil {
		panic("failed to migrate club")
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		panic("failed to migrate user")
	}

	return db
}

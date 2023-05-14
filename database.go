package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

type RecipeSuggestion struct {
	gorm.Model
	RecipeTitle string
	RecipeURL   string
	SuggestedOn time.Time
}

func GetDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot open database")
	}

	return db
}

func setupDatabase() {
	db := GetDatabaseConnection()

	_ = db.AutoMigrate(&RecipeSuggestion{})
}

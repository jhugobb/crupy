package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func handleDBCreate(c DnDCharacter) {

	db, err := gorm.Open("sqlite3", os.Getenv("DB_NAME"))
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&DnDCharacter{})

	db.Create(&c)
}

func handleDBFindAll(dndchars *[]DnDCharacter) *[]DnDCharacter {

	db, err := gorm.Open("sqlite3", os.Getenv("DB_NAME"))
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.Find(dndchars)

	return dndchars
}

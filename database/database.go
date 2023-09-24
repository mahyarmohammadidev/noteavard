package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"noteavard/note"
)

var DbInstance *gorm.DB
var err error

func ConnectToSqlite() {
	DbInstance, err = gorm.Open(sqlite.Open("note.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Could'nt establish a connection to sqlite database!")
	}

	log.Println("Connected to sqlite db")
}

func Migrate() {
	DbInstance.AutoMigrate(&note.Note{})
	log.Println("Migration is done.")
}

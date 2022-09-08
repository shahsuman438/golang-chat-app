package db

import (
	"log"
	"os"

	pg "github.com/go-pg/pg/v10"
)

func Connet() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "Suman@123",
		Addr:     "localhost:5432",
		Database: "postgres",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Connection to database successful.\n")
	CreateUserSchema(db)
	CreateChatSchema(db)
	return db
}

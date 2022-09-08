package db

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)



func (chat *Chat) Save(db *pg.DB) error {
	_, insertErr := db.Model(chat).Insert()
	if insertErr != nil {
		log.Printf("Error while Inserting data to table Chat, Reason: %v\n", insertErr)
		return insertErr
	}
	return nil
}

func CreateChatSchema(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&Chat{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table Chats, Reason: %v\n", createErr)
		return createErr
	}
	return nil
}

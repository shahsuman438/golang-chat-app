package controller

import (
	"log"
	"server/db"
	"server/services"
)

func SaveChat(chat *db.Chat) {
	err := services.SaveChat(chat)
	if err != nil {
		log.Println("Error while saving chat, Reason", err)
	}
}

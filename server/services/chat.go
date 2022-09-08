package services

import "server/db"

func SaveChat(chatData *db.Chat) error {
	pg_db := db.Connet()
	defer pg_db.Close()
	error := chatData.Save(pg_db)
	if error != nil {
		return error
	}
	return nil
}

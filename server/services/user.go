package services

import (
	db "server/db"
)

func CreateUser(userData *db.User) (*db.User, error) {
	pg_db := db.Connet()
	defer pg_db.Close()
	error := userData.Save(pg_db)
	if error != nil {
		return nil, error
	}
	return userData, nil
}

func GetUser(userData *db.User) (*db.User, error) {
	pg_db := db.Connet()
	defer pg_db.Close()
	result, getErr := userData.GetByEmail(pg_db)
	if getErr != nil {
		return nil, getErr
	}
	return result, nil
}

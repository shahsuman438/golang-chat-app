package db

import (
	"log"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func (user *User) Save(db *pg.DB) error {
	_, insertErr := db.Model(user).Insert()
	if insertErr != nil {
		log.Printf("Error while Inserting data to table User, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("user %s inserted successfully.\n", user.UserName)
	return nil
}

func CreateUserSchema(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.Model(&User{}).CreateTable(opts)
	if createErr != nil {
		log.Printf("Error while creating table User, Reason: %v\n", createErr)
		return createErr
	}
	return nil
}

func (user *User) GetByEmail(db *pg.DB) (*User, error) {
	err := db.Model(user).Column("id", "email", "user_name", "password").Where("email=?", user.Email).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

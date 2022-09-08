package db


type User struct {
	ID       int    `sql:"id,pk" json:"id"`
	Email    string `sql:"email,unique" json:"email"`
	UserName string `sql:"user_name" json:"userName"`
	Password string `sql:"password" json:"password"`
}

type Chat struct {
	Id          int    `sql:"id,pk" json:"id"`
	Source      string `sql:"source" json:"source"`
	Destination string `sql:"destination" json:"destination"`
	Message     string `sql:"Message" json:"Message"`
}
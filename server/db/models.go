package db

type User struct {
	ID       int    `pg:"id,pk" sql:"id,pk" json:"id"`
	Email    string `pg:"email,unique" sql:"email, unique" json:"email"`
	UserName string `pg:"user_name" sql:"user_name" json:"userName"`
	Password string `pg:"password" sql:"password" json:"password"`
}

type Chat struct {
	Id          int    `sql:"id,pk" json:"id"`
	Source      string `sql:"source" json:"source"`
	Destination string `sql:"destination" json:"destination"`
	Message     string `sql:"Message" json:"Message"`
}

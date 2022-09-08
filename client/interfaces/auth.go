package interfaces

type Login struct {
	Email    string
	Password string
}
type Register struct {
	Id       int
	Email    string
	Password string
	UserName string
}

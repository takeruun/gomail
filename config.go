package gomail

type Auth struct {
	Host     string // smtp host
	Email    string // email of smtp user
	Password string // password of smtp user
}

type From struct {
	Name  string // sender's name
	Email string // sender's email
}

type Config struct {
	Auth
	From
	Addr string // smt server
}

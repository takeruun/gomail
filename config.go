package gomail

type Config struct {
	Auth struct {
		Host     string
		Email    string
		Password string
	}
	From struct {
		Name  string
		Email string
	}
	Addr string
}

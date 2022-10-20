# Installation
```
go get github.com/takeruun/gomail
```

# Usage
## 1. Set your mail account and create gomail client
```go
config := &gomail.Config{
  Auth: gomail.Auth{
    Host:     config.Mail.Auth.Host,
    Email:    config.Mail.Auth.Email,
    Password: config.Mail.Auth.Password,
  },
  From: gomail.From{
    Name:  config.Mail.From.Name,
    Email: config.Mail.From.Email,
  },
  Addr: config.Mail.Addr,
}

mail := gomail.New(gomailConfig)
```

## 2. Send an email
```go
toEmail := "sender@email.com"
subject := "subject"
body := "body"

err := mail.Send(toEmail, subject, body)
```
# About
A utf8-capable mail sending library.

Currently, the only supported smtp server is a **gmail 2-step verification account**.

Accounts to be supported in the future
- office
- sendgrid

# Installation
```
go get github.com/takeruun/gomail
```

# Usage
## 1. Set your mail account and create gomail client
```go
config := &gomail.Config{
  Auth: gomail.Auth{
    Host:     "smtp-host",
    Email:    "auth-email",
    Password: "auth-password",
  },
  From: gomail.From{
    Name:  "sender name",
    Email: "sender email",
  },
  Addr: "smtp-address",
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

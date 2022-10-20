package gomail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

const (
	charactorLimitForOneLine = 78
)

type Mail struct {
	Auth smtp.Auth
	Addr string
	From mail.Address
}

func New(config *Config) (gomail *Mail) {
	gomail = &Mail{
		Auth: LoginAuth(config.Auth.Email, config.Auth.Password),
		Addr: config.Addr,
		From: mail.Address{Name: config.From.Name, Address: config.From.Email},
	}

	return
}

// 適切な長さにカットしCRLFを挿入
func cutAndAddCrlf(msg string) string {
	buffer := bytes.Buffer{}
	for k, c := range strings.Split(msg, "") {
		buffer.WriteString(c)
		if (k+1)%charactorLimitForOneLine == 0 {
			buffer.WriteString("\r\n")
		}
	}
	return buffer.String()
}

func makeMailBody(body string) string {
	encodedBody := base64.StdEncoding.EncodeToString([]byte(body))
	return cutAndAddCrlf(encodedBody)
}

// UTF8文字列を指定文字数で分割
func utf8Split(utf8string string, length int) []string {
	result := []string{}
	buffer := bytes.Buffer{}
	for k, c := range strings.Split(utf8string, "") {
		buffer.WriteString(c)
		if (k+1)%length == 0 {
			result = append(result, buffer.String())
			buffer.Reset()
		}
	}
	if buffer.Len() > 0 {
		result = append(result, buffer.String())
	}
	return result
}

// タイトルをMIMEエンコード
func encodeSubject(subject string) string {
	buffer := bytes.Buffer{}
	buffer.WriteString("Subject:")
	limit := charactorLimitForOneLine / 6 // Unicodeでは一文字が最大6バイトになるため
	for _, line := range utf8Split(subject, limit) {
		buffer.WriteString(" =?utf-8?B?")
		buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(line)))
		buffer.WriteString("?=\r\n")
	}
	return buffer.String()
}

// ヘッダを作る
func makeMailHeader(from, to, subject string) bytes.Buffer {
	header := bytes.Buffer{}
	header.WriteString("From: " + from + "\r\n")
	header.WriteString("To: " + to + "\r\n")
	header.WriteString(encodeSubject(subject))
	header.WriteString("MIME-Version: 1.0\r\n")
	header.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	header.WriteString("Content-Transfer-Encoding: base64\r\n")

	return header
}

func (mail *Mail) Send(to, subject, body string) error {
	mailHeader := makeMailHeader(mail.From.String(), to, subject)
	mailBody := makeMailBody(body)

	msg := mailHeader
	msg.WriteString(mailBody)

	if err := smtp.SendMail(mail.Addr, mail.Auth, mail.From.Address, []string{to}, msg.Bytes()); err != nil {
		return err
	}

	fmt.Print(msg, "\n")
	fmt.Print(body, "\n")

	return nil
}

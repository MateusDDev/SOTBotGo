package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ocorreu um erro carregando o arquivo .env")
	}

	token := os.Getenv("TOKEN")
	return token
}

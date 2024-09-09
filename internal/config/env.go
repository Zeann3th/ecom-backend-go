package config

import (
	"log"

	"github.com/joho/godotenv"
)

var Env map[string]string = initEnv()

func initEnv() map[string]string {
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return env
}

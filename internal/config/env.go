package config

import (
	"log"

	"github.com/joho/godotenv"
)

var Env map[string]string = initEnv()

func initEnv() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return env
}

func LoadEnv() error {
	err := godotenv.Load()
	return err
}

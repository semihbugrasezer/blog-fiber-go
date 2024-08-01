package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func GetPort() string {
    return os.Getenv("PORT")
}

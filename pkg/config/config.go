package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIURL            string
	Timezone          string
	DiscordWebhookURL string
}

func InitConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	fmt.Println(getEnv("API_URL", "https://api.example.com"))
	return Config{
		APIURL:            getEnv("API_URL", "https://api.example.com"),
		Timezone:          getEnv("TIMEZONE", "UTC"),
		DiscordWebhookURL: getEnv("DISCORD_WEBHOOK_URL", "https://discord.com/api/webhooks/1234567890/abcdefghijklmnopqrstuvwxyz"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenAIAPIKey string
	MemosBaseURL string
	MemosAPIKey  string
}

var ConfigInstance Config

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	openAIApiKey := os.Getenv("OPEN_AI_API_KEY")
	if openAIApiKey == "" {
		log.Fatal("OPEN_AI_API_KEY is not set in .env file")
	}
	memosBaseURL := os.Getenv("MEMOS_BASE_URL")
	if memosBaseURL == "" {
		log.Fatal("MEMOS_BASE_URL is not set in .env file")
	}
	memosAPIKey := os.Getenv("MEMOS_API_KEY")
	if memosAPIKey == "" {
		log.Fatal("MEMOS_API_KEY is not set in .env file")
	}

	ConfigInstance = Config{
		OpenAIAPIKey: openAIApiKey,
		MemosBaseURL: memosBaseURL,
		MemosAPIKey:  memosAPIKey,
	}

}

func GetConfig() Config {
	return ConfigInstance
}

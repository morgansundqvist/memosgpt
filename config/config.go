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
	memosBaseURL := os.Getenv("MEMOS_BASE_URL")
	memosAPIKey := os.Getenv("MEMOS_API_KEY")

	ConfigInstance = Config{
		OpenAIAPIKey: openAIApiKey,
		MemosBaseURL: memosBaseURL,
		MemosAPIKey:  memosAPIKey,
	}

}

func GetConfig() Config {
	return ConfigInstance
}

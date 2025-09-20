package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	MerchantID string
	PrivateKey string
	TestMode   bool
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		MerchantID: getEnv("SANTIMPAY_MERCHANT_ID", ""),
		PrivateKey: getEnv("SANTIMPAY_PRIVATE_KEY", ""),
		TestMode:   getEnvBool("SANTIMPAY_TEST_MODE", true),
	}

	if config.MerchantID == "" {
		log.Fatal("SANTIMPAY_MERCHANT_ID is required")
	}

	if config.PrivateKey == "" {
		log.Fatal("SANTIMPAY_PRIVATE_KEY is required")
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true" || value == "1"
	}
	return defaultValue
}
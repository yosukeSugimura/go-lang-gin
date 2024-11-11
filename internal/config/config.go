package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config構造体はアプリケーションの設定を保持します
type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadEnv は .env ファイルを読み込むための関数です
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}
}

// GetConfig はアプリケーションの設定を取得する関数です
func GetConfig() *Config {
	LoadEnv()
	return &Config{
		Port:       GetEnv("PORT", "8080"),
		DBHost:     GetEnv("DB_HOST", "localhost"),
		DBPort:     GetEnv("DB_PORT", "5432"),
		DBUser:     GetEnv("DB_USER", "user"),
		DBPassword: GetEnv("DB_PASSWORD", "password"),
		DBName:     GetEnv("DB_NAME", "mydatabase"),
	}
}

// getEnv は環境変数を取得し、デフォルト値を設定するヘルパー関数です
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

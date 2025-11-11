package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, akan menggunakan environment variables yang ada.")
	}

	host, err := cariEnv("DB_HOST")
	if err != nil {
		return "", err
	}
	port, err := cariEnv("DB_PORT")
	if err != nil {
		return "", err
	}
	user, err := cariEnv("DB_USER")
	if err != nil {
		return "", err
	}
	pass, err := cariEnv("DB_PASSWORD")
	if err != nil {
		return "", err
	}
	dbname, err := cariEnv("DB_NAME")
	if err != nil {
		return "", err
	}

	sslmode, err := cariEnv("DB_SSL_MODE")
	if err != nil {
		return "", err
	}

	konek := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, dbname, sslmode)

	return konek, nil
}

func cariEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return "", fmt.Errorf("environment variable %s tidak diatur atau kosong", key)
	}
	return value, nil
}

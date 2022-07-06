package infra

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName               string
	AppPort               string
	MongoDBURI            string
	MongoDBCollectionName string
	MongoDBDatabaseName   string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	config := &Config{
		AppName:               os.Getenv("APP_NAME"),
		AppPort:               os.Getenv("APP_PORT"),
		MongoDBURI:            os.Getenv("MONGODB_URI"),
		MongoDBCollectionName: os.Getenv("MONGODB_COLLECTION_NAME"),
		MongoDBDatabaseName:   os.Getenv("MONGODB_DATABASE_NAME"),
	}

	return config, err
}

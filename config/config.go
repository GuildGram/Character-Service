package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Name        string // Name of the app
	Environment string // Current environment

	RabbitMQConfig RabbitmqConfig
}

type RabbitmqConfig struct {
	Port     string
	Uri      string
	Username string
	Password string
	Queue    string
	Exchange string
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) LoadConfig() error {
	cfg.Name = os.Getenv("NAME")
	cfg.Environment = os.Getenv("ENVIRONMENT")

	if cfg.Environment != "develop" && cfg.Environment != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("[Config]: failed initializing .env with file")
			return err
		}
	}

	cfg.RabbitMQConfig = LoadRabbitmqConfig()

	return nil
}

func LoadRabbitmqConfig() RabbitmqConfig {
	return RabbitmqConfig{
		Port:     os.Getenv("RABBITMQ_PORT"),
		Uri:      os.Getenv("RABBITMQ_URI"),
		Username: os.Getenv("RABBITMQ_USERNAME"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
		Queue:    os.Getenv("RABBITMQ_QUEUE"),
		Exchange: os.Getenv("RABBITMQ_EXCHANGE"),
	}
}

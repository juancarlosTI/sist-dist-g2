package config

import (
	"fmt"
	"os"
)

type Config struct {
	// DB
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	RabbitURL  string
	HTTPPort   string
	// S3
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	UseSSL    bool
}

func Load() *Config {
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		RabbitURL:  os.Getenv("RABBIT_URL"),
		HTTPPort:   os.Getenv("HTTP_PORT"),
		Endpoint:   os.Getenv("MINIO_ENDPOINT"),
		AccessKey:  os.Getenv("MINIO_ACCESS_KEY"),
		SecretKey:  os.Getenv("MINIO_SECRET_KEY"),
		Bucket:     os.Getenv("MINIO_BUCKET"),
		UseSSL:     false,
	}
}

func (c *Config) PostgresDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

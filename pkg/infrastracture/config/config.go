package config

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	if err := godotenv.Load("pkg/infrastracture/config/.env"); err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

type MailConfig struct {
	HostName string
	Port     string
	Username string
	Password string
}

func Mail() *MailConfig {
	return &MailConfig{
		HostName: GetEnv("MAIL_HOST_NAME"),
		Port:     GetEnv("MAIL_PORT"),
		Username: GetEnv("MAIL_USER"),
		Password: GetEnv("MAIL_PASSWORD"),
	}
}

func Cors() cors.Config {
	return cors.Config{
		AllowOrigins: []string{
			GetEnv("CORS_AUTH_PORT"),
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"X-Secret-Key",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}
}

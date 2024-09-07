package config

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
)

type MailConfig struct {
	HostName string
	Port     string
	Username string
	Password string
}

func Mail() *MailConfig {
	return &MailConfig{
		HostName: os.Getenv("MAIL_HOST_NAME"),
		Port:     os.Getenv("MAIL_PORT"),
		Username: os.Getenv("MAIL_USER"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
}

func Cors() cors.Config {
	return cors.Config{
		AllowOrigins: []string{
			os.Getenv("CORS_AUTH_PORT"),
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

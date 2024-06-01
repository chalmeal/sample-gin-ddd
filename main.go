package main

import (
	"log"
	"os"
	"sample-gin-ddd/pkg/controller"
	"sample-gin-ddd/pkg/infrastracture/db"

	"github.com/joho/godotenv"
)

func init() {
	db.NewDB()
}

func main() {
	if err := godotenv.Load("pkg/infrastracture/config/.env"); err != nil {
		log.Fatal(err)
	}

	routes := controller.InitRouter()
	routes.Run(os.Getenv("SERVER_PORT"))
}

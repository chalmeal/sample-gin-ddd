package db

import (
	"log"
	"os"
	"sample-gin-ddd/pkg/model/fixture"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func NewDB() *gorm.DB {
	if err := godotenv.Load("pkg/infrastracture/config/.env"); err != nil {
		log.Fatal(err)
	}
	db, err = gorm.Open(mysql.Open(os.Getenv("DB_PATH")))
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("DATA_FIXTURE_ENABLE") != "NONE" {
		fix := fixture.NewDataFixtures(db)
		fix.DataFixture()
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}

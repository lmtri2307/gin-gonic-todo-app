package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL, isPresent := os.LookupEnv("DB_POSTGRES_URL")
	if !isPresent {
		log.Fatalf("Missing Postgres DB Url")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

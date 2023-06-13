package database

import (
	"fmt"
	"log"
	"manga-microservice/manga"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST_MANGA")
	dbPort := os.Getenv("DB_PORT_MANGA")
	dbUser := os.Getenv("DB_USER_MANGA")
	dbPass := os.Getenv("DB_PASS_MANGA")
	dbName := os.Getenv("DB_NAME_MANGA")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to manga database")
		return nil
	}
	fmt.Println("Succes to connect to manga database")

	db.AutoMigrate(&manga.Manga{})

	return db
}

// Here we have made connection with the postgres SQL server, the DB_URL is fetched from the .env file. The documentation for the code below can be found at "https://gorm.io/docs/connecting_to_the_database.html"
package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

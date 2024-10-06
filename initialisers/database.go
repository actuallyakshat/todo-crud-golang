package initialisers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	//Data Source Name

	dsn := os.Getenv("DB_URL")
	// Open the database connection
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Check for connection errors
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Database connection established successfully")
}

package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func COnnectDb() {
	var err error

	dsn := "host=localhost user=postgres password=harsh dbname=test port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(DB, "Failed to connect to database")
	} else {

		fmt.Println(DB, "Database connection established successfully")
	}
}

package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var connection string = "host=localhost user=postgres password=hH&qeV%y12 dbname=go_rest port=5432 sslmode=disable"
	db, connectionError := gorm.Open(postgres.Open(connection))

	if connectionError != nil {
		panic("Failed to connect to database!")
	}

	// db.AutoMigrate(
	// 	&models.Author{},
	// 	&models.Book{},
	// 	&models.Genre{},
	// 	&models.Publisher{},
	// )

	fmt.Println("Database connected successfully!")
	DB = db
}

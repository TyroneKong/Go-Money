package database

import (
	"finance/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	dns := fmt.Sprintf("%v:%v@/%v", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	connection, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Panic("could not connect to the database")
	} else {
		log.Println("connected to the database")
	}
	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Transaction{})
	connection.AutoMigrate(&models.Expense{})
	connection.AutoMigrate(&models.Budget{})

}

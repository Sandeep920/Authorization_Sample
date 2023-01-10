package injection

import (
	// Get mysql library
	"go-graphql-jwt/graph/model"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Create Database
func CreateDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Printf("[WARNING] %v", err)
		log.Printf("[WARNING] %v", ".env file doesn't exist, so please provide environment variables from some other way using .env.example as a reference ")
	}

	dsn := os.Getenv("DSN")

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		panic(err)
	}

	DB = db
}

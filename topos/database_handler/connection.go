package database_handler

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

// global variable to hold the db connection
var db *gorm.DB

func init() {
  e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

  // fetch credentials from .env file
	user     := os.Getenv("user")
	password := os.Getenv("password")
	dbname   := os.Getenv("dbname")
	host     := os.Getenv("host")
  port     := os.Getenv("port")


  // create uri
  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  // establish connection
	conn, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Print(err)
	}

  // assign connection to the global variable
	db = conn
}

func GetDB() *gorm.DB {
	return db
}

package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "password"
  dbname   = "topos"
)

func main() {
  // create a connection string with all the info
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  // validate arguements provided with the local postgres database
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  // create connection with the local postgres database
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  
  fmt.Println("Successfully connected!")
}

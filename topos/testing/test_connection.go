package main

import (
  "database/sql"
  "fmt"
  "os"
  "github.com/joho/godotenv"

  _ "github.com/lib/pq"
)

func main() {
  // fetch credentials from .env file
	user     := os.Getenv("user")
	password := os.Getenv("password")
	dbname   := os.Getenv("dbname")
	host     := os.Getenv("host")
  port     := os.Getenv("port")
  
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

  sqlStatement := `
    INSERT INTO ny_data (bin, cnstrct_yr, lstmoddate, lststatype, doitt_id, heightroof, feat_code,
                          groundelev, shape_area, shape_len, base_bbl, mpluto_bbl, geomsource)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
    RETURNING id`

  id := 0
  err = db.QueryRow(sqlStatement, 11231, 222222, "Jonathan", "Calhoun", 333533, 444444, 555555, 66666, 777777, 888888, "luigi", "waluigi", "mario").Scan(&id)
  if err != nil {
    panic(err)
  }

  fmt.Println("New record ID is:", id)
}

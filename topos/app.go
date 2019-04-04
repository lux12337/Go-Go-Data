package main

import (
  "database/sql"
	"encoding/json"
	"fmt"
	"log"
  "os"
	"net/http"
  "io/ioutil"
  "github.com/joho/godotenv"

  _ "github.com/lib/pq"
)

// struct to hold data from API
type Entry struct {
	Bin            string `json:"bin"`
	Cnstrct_yr     string `json:"cnstrct_yr"`
	Lstmoddate     string `json:"lstmoddate"`
	Lststatype     string `json:"lststatype"`
	Doitt_id       string `json:"doitt_id"`
	Heightroof     string `json:"heightroof"`
	Feat_code      string `json:"feat_code"`
	Groundelev     string `json:"groundelev"`
  Shape_area     string `json:"shape_area"`
  Shape_len      string `json:"shape_len"`
  Base_bbl       string `json:"base_bbl"`
  Mpluto_bbl     string `json:"mpluto_bbl"`
  Geomsource     string `json:"geomsource"`
}

func main() {
  //******************* Get data from API *******************//
  // API endpoint
  url := fmt.Sprintf("https://data.cityofnewyork.us/resource/hdxe-i756.json")

	// build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

  // using APP token for SODA
  req.Header.Set("X-App-Token", "sb6EoW4dhAxuEtIGPifEbV9Rl")
  client := &http.Client{}

  // send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

  // defer response close
	defer resp.Body.Close()

  // confirm we received an OK status
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

  // read the entire body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body:", err)
	}

  // create an empty instance of Entry struct
	var entries []Entry

  // unmarshal the JSON data into entries
  err = json.Unmarshal(body, &entries)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Println("Fetched data from API successfully")

  //******************* Store data in postgres database *******************//
  // create a connection string with all the credentials
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

  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
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

  // loop through the entries
	for _, ed := range entries {
    err = db.QueryRow(sqlStatement, ed.Bin, ed.Cnstrct_yr, ed.Lstmoddate, ed.Lststatype, ed.Doitt_id, ed.Heightroof,
      ed.Feat_code, ed.Groundelev, ed.Shape_area, ed.Shape_len, ed.Base_bbl, ed.Mpluto_bbl, ed.Geomsource).Scan(&id)
    if err != nil {
      panic(err)
    }

    fmt.Println("New record ID is:", id)
	}

  fmt.Println("Pushed all the data to postgres database successfully")
}

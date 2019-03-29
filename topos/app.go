package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "io/ioutil"
)

type Entry struct {
	Bin            string `json:"bin"`
	Cnstrct_yr     string `json:"cnstrct_yr"`
	Name           string `json:"name"`
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
  // API endpoint
  url := fmt.Sprintf("https://data.cityofnewyork.us/resource/k8ez-gyqp.json")

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

  fmt.Printf("%v\n", entries)
}

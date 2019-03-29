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

// the feed is the full JSON data structure
// this sets up the array of Entry types (defined above)
type Feed struct {
  Buildings []struct {
		Data Entry
	}
}

func main() {
  url := fmt.Sprintf("https://data.cityofnewyork.us/resource/k8ez-gyqp.json")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

  req.Header.Set("X-App-Token", "sb6EoW4dhAxuEtIGPifEbV9Rl")
  client := &http.Client{}

  // Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

  // Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
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

  // create an empty instance of Feed struct
	// this is what gets filled in when unmarshaling JSON
	var entries []Entry
  err = json.Unmarshal(body, &entries)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Printf("%v\n", entries)
}

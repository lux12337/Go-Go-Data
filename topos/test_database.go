package main

import (
	"topos/database_handler"
  "fmt"
)

func main() {
  var id uint = 1910
  data := database_handler.GetEntries(id)


  for _, ed := range data {
    fmt.Println(ed)
	}
}

package main

import (
	"topos/database_handler"
  "fmt"
)

func main() {
  var id uint = 10
  data := database_handler.GetEntry(id)
  fmt.Println(data)
}

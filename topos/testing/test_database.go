package main

import (
	"topos/database_handler"
  "fmt"
)

func main() {
  data := database_handler.GetAll()

  for _, ed := range data {
    fmt.Println(ed)
	}
}

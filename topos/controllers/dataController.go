package controllers

import (
	"net/http"
  "log"
	"topos/database_handler"
	u "topos/utility"
)

var GetAll = func(w http.ResponseWriter, r *http.Request) {

	data := database_handler.GetAll()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetAfterYear = func(w http.ResponseWriter, r *http.Request) {

	year, ok := r.URL.Query()["year"]
  if !ok || len(year[0]) < 1 {
        log.Println("Url Param 'year' is missing")
        return
      }

	data := database_handler.GetAfterYear(year)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetAvgHeightAfterYear = func(w http.ResponseWriter, r *http.Request) {

	year, ok := r.URL.Query()["year"]
  if !ok || len(year[0]) < 1 {
        log.Println("Url Param 'year' is missing")
        return
      }

	data := database_handler.GetAvgHeightAfterYear(year)
	resp := u.Message(true, "success")
	resp["averages"] = data
	u.Respond(w, resp)
}

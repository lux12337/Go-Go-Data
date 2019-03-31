package controllers

import (
	"net/http"
	"topos/database_handler"
	u "topos/utility"
)

var GetAll = func(w http.ResponseWriter, r *http.Request) {

	//id := r.Context().Value("user") . (uint)
	data := database_handler.GetAll()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

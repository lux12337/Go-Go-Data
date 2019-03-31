package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"topos/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/all", controllers.GetAll).Methods("GET")

	//router.NotFoundHandler = app.NotFoundHandler

	err := http.ListenAndServe(":8000", router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}

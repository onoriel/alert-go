package main

import (
	"log"
	"net/http"

	. "alerts/config"
	"alerts/controller"

	. "alerts/repository"

	"github.com/gorilla/mux"
)

var config = Config{}
var repository = AlertsRepository{}

// Load app configuration from config.yaml, and establish a connection to DB
func init() {

	config.Read()

	repository.Server = config.Server
	repository.DatabaseName = config.Database

	repository.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()

	alertController := controller.NewAlertController(repository)
	r.HandleFunc("/alerts", alertController.AllAlertsEndPoint).Methods("GET")
	r.HandleFunc("/alerts", alertController.CreateAlertEndPoint).Methods("POST")
	r.HandleFunc("/alerts", alertController.UpdateAlertEndPoint).Methods("PUT")
	r.HandleFunc("/alerts", alertController.DeleteAlertEndPoint).Methods("DELETE")
	r.HandleFunc("/alerts/{id}", alertController.FindAlertEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"golang-crud/config"
	"golang-crud/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

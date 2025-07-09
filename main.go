package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TiagoAmaralFerreira/go-expert-cloud-run/handlers"
)

func main() {
	http.HandleFunc("/weather/", handlers.WeatherHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

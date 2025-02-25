package main

import (
	"fmt"
	"go-web-app/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/health", handlers.HealthHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

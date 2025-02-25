package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler returns "Hello, World!"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// HealthHandler returns "OK" for health check
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

package handler

import (
	"fmt"
	"net/http"
)

// RootHandler handles the root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Server is running")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

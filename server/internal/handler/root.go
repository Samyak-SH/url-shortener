package handler

import (
	"fmt"
	"net/http"
)

// RootHandler handles the root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	if r.URL.Path == "/test" {
		fmt.Fprintf(w, "Server is running")
		return
	} else if r.URL.Path != "/" {
		GetUrl(&w, r)
	} else {
		fmt.Fprintf(w, "bruh")
	}
}

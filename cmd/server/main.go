package main

import (
	"fmt"
	"net/http"

	"url-shortener/internal/handler"
)

const PORT = 8080

func main() {
	fmt.Printf("Server running on http://localhost:%d\n", PORT)

	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/create", handler.CreateHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

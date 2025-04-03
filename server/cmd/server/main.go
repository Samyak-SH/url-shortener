package main

import (
	"fmt"
	"log"
	"net/http"

	"url-shortener/internal/handler"
)

const PORT = 8080

func main() {
	fmt.Printf("Server running on http://localhost:%d\n", PORT)

	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/url", handler.UrlHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

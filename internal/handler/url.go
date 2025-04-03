package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/internal/shortener"
	"url-shortener/internal/storage"
)

// CreateHandler handles URL shortening requests
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		URL string `json:"original_url"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode data: %s", err), http.StatusBadRequest)
		return
	}

	shortURL := shortener.GenerateShortUrl(request.URL)
	storage.AddToDb(request.URL, shortURL)

	response := map[string]string{"short_url": shortURL}
	json.NewEncoder(w).Encode(response)
}

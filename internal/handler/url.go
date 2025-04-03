package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/internal/shortener"
	"url-shortener/internal/storage"
)

func createUrl(w *http.ResponseWriter, r *http.Request) {

	// fmt.Println(r.Header)
	parseFormErr := r.ParseForm() //populates r.Form
	if parseFormErr != nil {
		http.Error(*w, fmt.Sprintf("Failed to parse form : %s", parseFormErr.Error()), http.StatusBadRequest)
		fmt.Println("failed to parse form")
	}

	var original_url string = r.FormValue("original_url")
	if original_url == "" {
		http.Error(*w, fmt.Sprintf("Invalid URL: %s", original_url), http.StatusBadRequest)
		return
	}

	shortURL := shortener.GenerateShortUrl(original_url)
	storage.AddToDb(original_url, shortURL)

	response := map[string]string{"short_url": shortURL}
	json.NewEncoder(*w).Encode(response)
}

func GetUrl(w *http.ResponseWriter, r *http.Request) {

	var short_url string = r.URL.Path[1:] //remove the leading '/'

	original_url, err := storage.GetFromDb(short_url)

	if err != nil {
		fmt.Println("Failed to fetch from DB :", err.Error())
		http.Error(*w, fmt.Sprintf("Failed to fetch from DB : %s", err.Error()), http.StatusAccepted)
		return
	}

	type Response struct {
		Original_url string
	}

	response := Response{
		Original_url: original_url,
	}
	fmt.Println(response)

	json.NewEncoder(*w).Encode(response)

}

// CreateHandler handles URL shortening requests
func UrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createUrl(&w, r)
		return
	} else {
		http.Error(w, "Method not allowed", http.StatusBadGateway)
	}

}

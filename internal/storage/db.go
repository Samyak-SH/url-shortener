package storage

import "errors"

// URL struct to store original URLs
type URL struct {
	OriginalURL string `json:"original_url"`
}

// In-memory database
var data = make(map[string]URL)

// AddToDb stores a URL mapping
func AddToDb(ogURL, shortURL string) {
	data[shortURL] = URL{OriginalURL: ogURL}
}

// GetUrl retrieves the original URL
func GetUrl(shortURL string) (string, error) {
	urlObj, exists := data[shortURL]
	if exists {
		return urlObj.OriginalURL, nil
	}
	return "", errors.New("URL not found")
}

package shortener

import (
	"crypto/sha256"
	"encoding/hex"
)

// GenerateShortUrl creates a short URL using SHA256
func GenerateShortUrl(ogURL string) string {
	hasher := sha256.New()
	hasher.Write([]byte(ogURL))
	hashedData := hasher.Sum(nil)
	finalHashedString := hex.EncodeToString(hashedData)

	return finalHashedString[:5] // Get only the first 5 characters
}

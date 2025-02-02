package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"github.com/gocroot/config"
)

// Fungsi untuk menghasilkan token CSRF
func GenerateCSRFToken() (string, error) {
	// Membuat buffer untuk menyimpan token
	b := make([]byte, 32) // Token 32 byte
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Encode dalam format base64
	return base64.StdEncoding.EncodeToString(b), nil
}

// Handler untuk menangani request CSRF Token
func CSRFToken(w http.ResponseWriter, r *http.Request) {
	token := config.GenerateCSRFToken() // Menghasilkan token CSRF
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"csrf_token": "` + token + `"}`)) // Mengirimkan token dalam response JSON
}
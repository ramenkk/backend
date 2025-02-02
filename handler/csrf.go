package handler

import (
	"net/http"
	"github.com/gocroot/config"
)

// Handler untuk menangani request CSRF Token
func CSRFToken(w http.ResponseWriter, r *http.Request) {
	// Menghasilkan token CSRF dengan fungsi dari config
	token := config.GenerateCSRFToken() 
	if token == "" {
		http.Error(w, "Failed to generate CSRF token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"csrf_token": "` + token + `"}`)) // Mengirimkan token dalam response JSON
}

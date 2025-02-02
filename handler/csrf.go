package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/csrf"
)

// CSRFToken mengembalikan token CSRF ke frontend
func CSRFToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-CSRF-Token", token) // Tambahkan token ke header

	json.NewEncoder(w).Encode(map[string]string{
		"csrf_token": token,
	})
}

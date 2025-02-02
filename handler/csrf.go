package handler

import (
    "net/http"
    "github.com/gorilla/csrf"
)

// Handler untuk menangani request CSRF Token
func CSRFToken(w http.ResponseWriter, r *http.Request) {
    token := csrf.Token(r) // Token yang dikelola oleh gorilla/csrf
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"csrf_token": "` + token + `"}`)) // Mengirimkan token dalam response JSON
}

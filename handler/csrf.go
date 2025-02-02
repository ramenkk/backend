package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/csrf"
)

// CSRFToken mengembalikan token CSRF ke frontend
func CSRFToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"csrf_token": csrf.Token(r),
	})
}

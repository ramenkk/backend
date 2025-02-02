package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

// CSRFToken mengembalikan token CSRF ke frontend
func CSRFToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	log.Println("Generated CSRF Token:", token) // Log token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"csrftoken": token})
}

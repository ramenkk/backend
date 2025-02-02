package middleware

import (
	"log"
	"net/http"
	"github.com/gorilla/csrf"
)

var csrfKey = []byte("super-secret-32-byte-key") // Ganti dengan key yang lebih aman

// Middleware CSRF Protection
var CSRF = csrf.Protect(
	csrfKey,
	csrf.Secure(false), // true jika menggunakan HTTPS
	csrf.HttpOnly(true),
	csrf.Path("/"),
	csrf.CookieName("csrftoken"), // Pastikan cookie disimpan dengan benar
)

// Middleware Handler untuk validasi CSRF pada request
func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-CSRF-Token")
		log.Println("Received CSRF Token:", token)
		next.ServeHTTP(w, r)
	})
}

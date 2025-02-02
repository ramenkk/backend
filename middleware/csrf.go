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
	csrf.Secure(false), // true jika pakai HTTPS
	csrf.HttpOnly(true),
	csrf.Path("/"),
	csrf.CookieName("csrftoken"), // Pastikan cookie sesuai
	csrf.RequestHeader("X-CSRF-Token"),
)

func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := csrf.Token(r)
		log.Println("Middleware Generated CSRF Token:", token) // Tambahkan log

		w.Header().Set("X-CSRF-Token", token) // Debugging
		next.ServeHTTP(w, r)
	})
}

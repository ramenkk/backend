package middleware

import (
	"github.com/gorilla/csrf"
	"net/http"
)

var csrfKey = []byte("super-secret-32-byte-key") // Ganti dengan key yang lebih aman

// Middleware CSRF Protection
var CSRF = csrf.Protect(
	csrfKey,
	csrf.Secure(false), // Ubah ke `true` jika menggunakan HTTPS
	csrf.HttpOnly(true),
	csrf.Path("/"),
)

// Middleware Handler untuk validasi CSRF pada request
func CSRFMiddleware(next http.Handler) http.Handler {
	return CSRF(next)
}

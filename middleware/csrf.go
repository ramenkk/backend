package middleware

import (
	"net/http"
	"github.com/gorilla/csrf"
)

var csrfKey = []byte("32-byte-long-auth-key") // Gantilah dengan kunci yang lebih aman

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

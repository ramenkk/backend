package middleware

import (
	"github.com/gorilla/csrf"
	"net/http"
)

// Middleware CSRF Protection
func CSRFMiddleware(next http.Handler) http.Handler {
	secretKey := "your-secret-key" // Gunakan key yang lebih aman di env variable

	csrfMiddleware := csrf.Protect(
		[]byte(secretKey),
		csrf.Secure(true),   // Hanya aktif jika menggunakan HTTPS
		csrf.HttpOnly(true), // Cookie hanya bisa diakses oleh server
		csrf.Path("/"),      // Berlaku untuk semua endpoint
	)

	return csrfMiddleware(next)
}

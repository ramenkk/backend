package middleware

import (
	"log"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gorilla/csrf"
)

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

// Middleware untuk menangani validasi CSRF token
func CSRFValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-CSRF-Token")
		if token == "" || !config.IsValidCSRFToken(token) {
			http.Error(w, "Forbidden - CSRF token is invalid", http.StatusForbidden)
			log.Printf("Invalid CSRF token received: %s", token) // Tambahkan log di sini
			return
		}

		next.ServeHTTP(w, r)

		config.RemoveCSRFToken(token)
	})
}

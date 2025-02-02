package middleware

import (
	"net/http"
	"github.com/gorilla/csrf"
)

/// Middleware CSRF Protection
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
        // `csrf.Protect` akan memvalidasi token CSRF untuk kita
        csrf.Protect([]byte("your-secret-key"))(next).ServeHTTP(w, r)
    })
}

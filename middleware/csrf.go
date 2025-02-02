package middleware

import (
	"net/http"
	"github.com/gorilla/csrf"
	"github.com/gocroot/config"
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
		// Ambil token CSRF dari header request
		token := r.Header.Get("X-CSRF-Token")
		if token == "" || !config.IsValidCSRFToken(token) {
			http.Error(w, "Forbidden - CSRF token is invalid", http.StatusForbidden)
			return
		}

		// Jika token valid, lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)

		// Hapus token setelah validasi (jika diperlukan untuk single-use token)
		config.RemoveCSRFToken(token)
	})
}

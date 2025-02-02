package gocroot

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocroot/middleware"
	"github.com/gocroot/route"
	"github.com/gorilla/csrf"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port jika tidak ada di env
	}

	// Middleware CSRF Protection
	csrfMiddleware := csrf.Protect(
		[]byte("your-secret-key"), // Ganti dengan key yang aman
		csrf.Secure(true),         // Aktifkan hanya jika menggunakan HTTPS
		csrf.HttpOnly(true),       // Cookie hanya bisa diakses oleh server
		csrf.Path("/"),            // Berlaku untuk semua endpoint
	)

	// Gunakan middleware CSRF & Auth pada semua request
	mux := http.NewServeMux()
	mux.HandleFunc("/", route.URL)

	// Jalankan server dengan middleware
	fmt.Println("Server berjalan di port", port)
	log.Fatal(http.ListenAndServe(":"+port, csrfMiddleware(middleware.AuthMiddleware(mux))))
}

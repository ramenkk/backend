package config

import (
	"net/http"
	"strings"
)

var AllowedOrigins = []string{
	"http://127.0.0.1:5500",
	"http://127.0.0.1:5501",
	"http://127.0.0.1:5500",
	"http://127.0.0.1:5502",
	"https://ramenkk.github.io",
	"https://menu.github.io",
	"https://ramenkk.github.io/kasir/",
	"https://irgifauzi.github.io",
	"https://irgifauzi.github.io/P2_kasir/",
}

// test
var AllowedHeaders = []string{
	"Origin",
	"Content-Type",
	"X-CSRF-Token",
	"Accept",
	"Authorization",
	"Access-Control-Request-Headers",
	"Token",
	"Login",
	"Access-Control-Allow-Origin",
	"Bearer",
	"X-Requested-With",
}

func SetAccessControlHeaders(w http.ResponseWriter, r *http.Request) bool {
	origin := r.Header.Get("Origin")

	allowedOrigin := false
	for _, o := range AllowedOrigins {
		if o == origin {
			allowedOrigin = true
			break
		}
	}
	if !allowedOrigin {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return false
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(AllowedHeaders, ", "))
	w.Header().Set("Access-Control-Allow-Origin", origin)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	}

	return false
}

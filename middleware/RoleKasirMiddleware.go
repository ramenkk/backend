package middleware

import (
	"net/http"
	"strings"
)

func RoleKasirMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims, err := validateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != "kasir" {
			http.Error(w, "Unauthorized: role kasir required", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

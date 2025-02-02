package config

import (
	"errors"
	"sync"
)

var csrfTokenStore = struct {
	sync.RWMutex
	tokens map[string]bool
}{tokens: make(map[string]bool)}

// GenerateCSRFToken menghasilkan token CSRF baru dan menyimpannya
func GenerateCSRFToken() string {
	// Token dapat dihasilkan dengan menggunakan random bytes atau cara lain
	token, _ := GenerateJWT("dummy", "dummy") // Anda dapat menyesuaikan dengan token yang lebih aman
	csrfTokenStore.Lock()
	csrfTokenStore.tokens[token] = true
	csrfTokenStore.Unlock()
	return token
}

// IsValidCSRFToken memvalidasi apakah token CSRF ada dan valid
func IsValidCSRFToken(token string) bool {
	csrfTokenStore.RLock()
	defer csrfTokenStore.RUnlock()
	_, exists := csrfTokenStore.tokens[token]
	return exists
}

// RemoveCSRFToken menghapus token CSRF dari penyimpanan setelah digunakan
func RemoveCSRFToken(token string) error {
	csrfTokenStore.Lock()
	defer csrfTokenStore.Unlock()
	if _, exists := csrfTokenStore.tokens[token]; !exists {
		return errors.New("token not found")
	}
	delete(csrfTokenStore.tokens, token)
	return nil
}

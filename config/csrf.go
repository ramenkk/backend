package config

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
)

var csrfTokenStore = struct {
	sync.RWMutex
	tokens map[string]bool
}{tokens: make(map[string]bool)}

// GenerateCSRFToken menghasilkan token CSRF baru dan menyimpannya
func GenerateCSRFToken() string {
	// Membuat buffer untuk menyimpan token
	b := make([]byte, 32) // Token 32 byte
	_, err := rand.Read(b)
	if err != nil {
		return "" // Kembalikan token kosong jika error
	}

	// Encode dalam format base64
	token := base64.StdEncoding.EncodeToString(b)

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

package helper

import (
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("this_is_a_32_byte_key_for_paseto") // Harus 32 byte

// HashPassword membuat hash dari password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPasswordHash memeriksa kesamaan password dengan hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken membuat token PASETO
func GenerateToken(username string, role string) (string, error) {
	token := paseto.NewV2()
	now := time.Now()
	expiration := now.Add(1 * time.Hour)

	payload := map[string]interface{}{
		"username": username,
		"role":     role,
		"exp":      expiration,
	}

	return token.Encrypt(secretKey, payload, nil)
}

// VerifyToken memverifikasi token PASETO
func VerifyToken(token string) (map[string]interface{}, error) {
	var payload map[string]interface{}
	err := paseto.NewV2().Decrypt(token, secretKey, &payload, nil)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

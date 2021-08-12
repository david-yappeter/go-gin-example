package tools

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hashed)
}

func CompareHash(hashed string, realPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(realPass))
}

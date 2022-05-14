package security

import "golang.org/x/crypto/bcrypt"

// Converts received string password to a hashed one and returns it
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Verifies if the hashed password is the same as the password
func VerifyPassword(hashPassword []byte, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(stringPassword))
}

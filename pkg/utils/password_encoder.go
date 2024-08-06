package utils

import "golang.org/x/crypto/bcrypt"

func Encode(plaintext string) (string, error) {
	cipherbytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	return string(cipherbytes), err
}

func Matches(ciphertext, plaintext string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ciphertext), []byte(plaintext))
	return err == nil
}

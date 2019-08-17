package password

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenSalt generate salt byte values with given length
func GenSalt(length int) []byte {
	salt := make([]byte, length)
	for i := range salt {
		r := rand.Intn(len(letters))
		salt[i] = letters[r]
	}
	return salt
}

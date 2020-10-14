package cipher

import (
	"crypto/rand"
	"io"
	"testing"
)

func TestAESCGMEncrypt(t *testing.T) {
	plaintext := []byte("Golang Programs")
	key := []byte("12345678901234567890123456789012")
	nonce := make([]byte, 12)
	additionalData := []byte("abcdefghijklmnopqrstuvwxyz")
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		t.Fatal(err)
	}

	// encrypt
	cipherText, err := AESCGMEncrypt(plaintext, nonce, key, additionalData)
	if err != nil {
		t.Fatal(err)
	}

	// decrypt
	plain, err := AESCGMDecrypt(cipherText, nonce, key, additionalData)
	if err != nil {
		t.Fatal(err)
	}
	if string(plain) != string(plaintext) {
		t.Fatalf("exp : %s, got: %s", string(plaintext), string(plain))
	}
}

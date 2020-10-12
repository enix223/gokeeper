package cipher_test

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/enix223/gokeeper/cipher"
)

func ExampleAESCGMEncrypt() {
	plaintext := []byte("Golang Programs")
	key := []byte("12345678901234567890123456789012")
	nonce := make([]byte, 12)
	additionalData := []byte("abcdefghijklmnopqrstuvwxyz")
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// encrypt
	cipherText, err := cipher.AESCGMEncrypt(plaintext, nonce, key, additionalData)
	if err != nil {
		panic(err.Error())
	}

	// decrypt
	plain, err := cipher.AESCGMDecrypt(cipherText, nonce, key, additionalData)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("plain: %s", string(plain))
	// Output: plain: Golang Programs
}

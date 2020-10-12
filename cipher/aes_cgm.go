package cipher

import (
	"crypto/aes"
	"crypto/cipher"
)

// AESCGMEncrypt AES CGM encrypt
// AEAD_AES_256_GCM
func AESCGMEncrypt(plainText, nonce, key, additionalData []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	cipherText := aesgcm.Seal(nil, nonce, plainText, additionalData)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

// AESCGMDecrypt AES CGM decrypt
// AEAD_AES_256_GCM
func AESCGMDecrypt(cipherText, nonce, key, additionalData []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainText, err := aesgcm.Open(nil, nonce, cipherText, additionalData)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

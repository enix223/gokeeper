package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// AESCFBEncrypt AES-CFB encrypt
func AESCFBEncrypt(plaintext []byte, key []byte, paddingType ...string) (cipherText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrInvalidKeyLen
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case PaddingTypeZero:
			plaintext = ZeroPadding(plaintext, aes.BlockSize)
		case PaddingTypePKCS5:
			plaintext = PKCS5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	}

	cipherText = make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText[aes.BlockSize:],
		plaintext)
	return cipherText, nil
}

// AESCFBDecrypt AES-CBF decrypt
func AESCFBDecrypt(cipherText []byte, key []byte, paddingType ...string) (plaintext []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrInvalidKeyLen
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(cipherText, cipherText)
	if int(cipherText[len(cipherText)-1]) > len(cipherText) {
		return nil, errors.New("aes decrypt failed")
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case PaddingTypeZero:
			plaintext = ZeroTrimming(cipherText)
		case PaddingTypePKCS5:
			plaintext = PKCS5Trimming(cipherText)
		}
	} else {
		plaintext = PKCS5Trimming(cipherText)
	}
	return plaintext, nil
}

package cipher

import "bytes"

const (
	// PaddingTypeZero zero padding
	PaddingTypeZero = "ZeroPadding"
	// PaddingTypePKCS5 PKCS5 padding
	PaddingTypePKCS5 = "PKCS5Padding"
	// PaddingTypePKCS7 PKCS7 padding
	PaddingTypePKCS7 = "PKCS7Padding"
)

// PKCS5Padding PKCS5 Padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	return PKCS7Padding(ciphertext, blockSize)
}

// PKCS5Unpadding PKCS5 unpadding
func PKCS5Unpadding(encrypt []byte) []byte {
	return PKCS7Unpadding(encrypt)
}

// PKCS7Padding PKCS7 Padding
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7Unpadding PKCS7 unpadding
func PKCS7Unpadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// ZeroPadding zero padding
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

// ZeroUnpadding zero unpadding
func ZeroUnpadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

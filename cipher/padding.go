package cipher

import "bytes"

const (
	// PaddingTypeZero zero padding
	PaddingTypeZero = "ZeroPadding"
	// PaddingTypePKCS5 PKCS5 padding
	PaddingTypePKCS5 = "PKCS5Padding"
)

// PKCS5Padding PKCS5 Padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5Trimming PKCS5 trimming
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// ZeroPadding zero padding
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

// ZeroTrimming zero trimming
func ZeroTrimming(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

package cipher

import (
	"crypto/aes"
	"crypto/cipher"
)

// AESCBCEncrypt AES CBC encrypt
func AESCBCEncrypt(plainText []byte, key, iv []byte, paddingType ...string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(plainText) == 0 {
		return nil, ErrInvalidPlainText
	}
	ecb := cipher.NewCBCEncrypter(block, iv)
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case PaddingTypeZero:
			plainText = ZeroPadding(plainText, block.BlockSize())
		case PaddingTypePKCS5:
			plainText = PKCS5Padding(plainText, block.BlockSize())
		case PaddingTypePKCS7:
			plainText = PKCS7Padding(plainText, block.BlockSize())
		}
	} else {
		plainText = PKCS5Padding(plainText, block.BlockSize())
	}
	crypted := make([]byte, len(plainText))
	ecb.CryptBlocks(crypted, plainText)

	return crypted, nil
}

// AESCBCDecrypt AES CBC decrypt
func AESCBCDecrypt(cipherText []byte, key, iv []byte, paddingType ...string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(cipherText) == 0 {
		return nil, ErrInvalidPlainText
	}
	ecb := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(cipherText))
	ecb.CryptBlocks(plaintext, cipherText)
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case PaddingTypeZero:
			plaintext = ZeroUnpadding(plaintext)
		case PaddingTypePKCS5:
			plaintext = PKCS5Unpadding(plaintext)
		case PaddingTypePKCS7:
			plaintext = PKCS7Unpadding(plaintext)
		}
	} else {
		plaintext = PKCS5Unpadding(plaintext)
	}
	return plaintext, nil
}

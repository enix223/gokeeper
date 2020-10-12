package cipher

import (
	"crypto/aes"
	"crypto/cipher"
)

// AESCBCEncrypt AES CBC encrypt
func AESCBCEncrypt(plainText string, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if plainText == "" {
		return nil, ErrInvalidPlainText
	}
	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(plainText)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted, nil
}

// AESCBCDecrypt AES CBC decrypt
func AESCBCDecrypt(crypt []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(crypt) == 0 {
		return nil, ErrInvalidPlainText
	}
	ecb := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted), nil
}

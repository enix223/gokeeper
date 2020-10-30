package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// newECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func newECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func newECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// AESECBEncrypt AES-ECB encrypt
func AESECBEncrypt(plaintext []byte, key []byte, paddingType ...string) (ciphertext []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrInvalidKeyLen
	}

	if len(paddingType) > 0 {
		switch paddingType[0] {
		case PaddingTypeZero:
			plaintext = ZeroPadding(plaintext, aes.BlockSize)
		case PaddingTypePKCS5:
			plaintext = PKCS5Padding(plaintext, aes.BlockSize)
		case PaddingTypePKCS7:
			plaintext = PKCS7Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	}
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext = make([]byte, len(plaintext))
	newECBEncrypter(block).CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

// AESECBDecrypt AES-ECB decrypt
func AESECBDecrypt(ciphertext []byte, key []byte, paddingType ...string) (plaintext []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrInvalidKeyLen
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, ErrCipherTooShort
	}
	// ECB mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, ErrCipherInvalidLength
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	newECBDecrypter(block).CryptBlocks(ciphertext, ciphertext)
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case PaddingTypeZero:
			plaintext = ZeroUnpadding(ciphertext)
		case PaddingTypePKCS5:
			plaintext = PKCS5Unpadding(ciphertext)
		case PaddingTypePKCS7:
			plaintext = PKCS7Unpadding(ciphertext)
		}
	} else {
		plaintext = PKCS5Unpadding(ciphertext)
	}
	return plaintext, nil
}

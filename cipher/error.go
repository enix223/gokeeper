package cipher

import "errors"

var (
	// ErrInvalidKey invalid key
	ErrInvalidKey = errors.New("AES: invalid key")
	// ErrInvalidPlainText invalid plain text
	ErrInvalidPlainText = errors.New("plain should not be empty")
	// ErrInvalidKeyLen invalid key len
	ErrInvalidKeyLen = errors.New("key size must be 16, 24 or 32")
	// ErrCipherTooShort cipher too short
	ErrCipherTooShort = errors.New("ciphertext too short")
	// ErrCipherInvalidLength cipher length invalid
	ErrCipherInvalidLength = errors.New("ciphertext is not a multiple of the block size")
)

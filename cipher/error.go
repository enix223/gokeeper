package cipher

import "errors"

var (
	// ErrInvalidKey invalid key
	ErrInvalidKey = errors.New("AES: invalid key")
	// ErrInvalidPlainText invalid plain text
	ErrInvalidPlainText = errors.New("plain should not be empty")
)

package password

import "fmt"

// HasherFactory factory to create PasswordHasher interface
type HasherFactory struct {
}

// HasherType password hasher
type HasherType uint8

const (
	// HasherPBKDF2 PBKDF2 algorithm
	HasherPBKDF2 HasherType = iota
)

// GetPasswordHasher create password hasher
func (f *HasherFactory) GetPasswordHasher(hasherType HasherType) (Hasher, error) {
	switch hasherType {
	case HasherPBKDF2:
		return &PBKDF2PasswordHasher{}, nil
	default:
		return nil, fmt.Errorf("Unknown password hasher type: %d", hasherType)
	}
}

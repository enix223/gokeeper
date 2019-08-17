package password

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const pbkdf2Algorithm = "pbkdf2_sha256"

// PBKDF2PasswordHasher Password hasher with PBKDF2
type PBKDF2PasswordHasher struct {
}

// MakePassword create password hash
func (h *PBKDF2PasswordHasher) MakePassword(password, salt string, iterations int) string {
	// generate salt
	if iterations <= 0 {
		iterations = 4096
	}

	saltByte := []byte(salt)
	if len(salt) == 0 {
		saltByte = GenSalt(12)
	}

	// generate pbkdf2
	dk := pbkdf2.Key([]byte(password), saltByte, iterations, 32, sha256.New)
	encoded := base64.StdEncoding.EncodeToString(dk)
	return fmt.Sprintf("%s$%d$%s$%s", pbkdf2Algorithm, iterations, string(saltByte), encoded)
}

// ValidatePassword validate the password with given encoded password
func (h *PBKDF2PasswordHasher) ValidatePassword(encodedPassword string, password string) bool {
	comps := strings.Split(encodedPassword, "$")
	if len(comps) != 4 {
		return false
	}

	algorithm, iterationStr, salt := comps[0], comps[1], comps[2]
	if algorithm != pbkdf2Algorithm {
		return false
	}

	var iterations int
	var err error
	if iterations, err = strconv.Atoi(iterationStr); err != nil {
		return false
	}

	password2 := h.MakePassword(password, salt, iterations)
	return password2 == encodedPassword
}

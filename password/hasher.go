package password

// Hasher create password hash
type Hasher interface {

	// MakePassword create password hash
	MakePassword(password, salt string, iterations int) string

	// ValidatePassword validate the password with given encoded password
	ValidatePassword(encodedPassword string, password string) bool
}

package cipher_test

import (
	"encoding/base64"
	"fmt"

	"github.com/enix223/gokeeper/cipher"
)

func ExampleAESCBCEncrypt() {
	var plainText = "Golang Programs"
	var initialVector = "1234567890123456"
	var passphrase = "Impassphrasegood"

	// encrypt
	encryptedData := cipher.AESCBCEncrypt(plainText, []byte(passphrase), []byte(initialVector))
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)

	// decrypt
	encryptedData, _ = base64.StdEncoding.DecodeString(encryptedString)
	decryptedText := cipher.AESCBCDecrypt(encryptedData, []byte(passphrase), []byte(initialVector))
	fmt.Printf("plain: %s", string(decryptedText))
	// Output: plain: Golang Programs
}

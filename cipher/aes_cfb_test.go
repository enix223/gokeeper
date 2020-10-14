package cipher

import (
	"encoding/base64"
	"testing"
)

// key长度 16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
func TestAESCFBDecrypt(t *testing.T) {
	key := []byte("nanjishidu170501")
	plainText := []byte("aes-20170501-30-1000")
	paddings := [][]string{
		{PaddingTypeZero},
		{PaddingTypePKCS5},
		{},
	}
	for _, padding := range paddings {

		se, err := AESCFBEncrypt(plainText, key, padding...)
		if err != nil {
			t.Fatal(err)
		}
		s := base64.StdEncoding.EncodeToString(se)
		sed, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			t.Fatal(err)
		}
		sd, err := AESCFBDecrypt(sed, key, padding...)
		if err != nil {
			t.Fatal(err)
		}
		if string(sd) != string(plainText) {
			t.Log(string(sd))
			t.FailNow()
		}
	}
}

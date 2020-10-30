package cipher

import (
	"encoding/base64"
	"testing"
)

func TestAESECBDecrypt(t *testing.T) {
	key := []byte("nanjishidu170502")
	plainText := []byte("aes-20171125-30-1002")
	paddings := [][]string{
		{PaddingTypeZero},
		{PaddingTypePKCS5},
		{PaddingTypePKCS7},
		{},
	}
	for _, padding := range paddings {
		se, err := AESECBEncrypt(plainText, key, padding...)
		if err != nil {
			t.Fatal(err)
		}
		s := base64.StdEncoding.EncodeToString(se)
		sed, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			t.Fatal(err)
		}
		sd, err := AESECBDecrypt(sed, key, padding...)
		if err != nil {
			t.Fatal(err)
		}
		if string(sd) != string(plainText) {
			t.Log(string(sd))
			t.FailNow()
		}
	}
}

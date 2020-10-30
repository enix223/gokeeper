package cipher

import (
	"encoding/base64"
	"testing"
)

func TestAESCBCDecrypt(t *testing.T) {
	key := []byte("nanjishidu170502")
	initialVector := []byte("1234567890123456")
	paddings := [][]string{
		{PaddingTypeZero},
		{PaddingTypePKCS5},
		{PaddingTypePKCS7},
		{},
	}

	for _, padding := range paddings {
		se, err := AESCBCEncrypt([]byte("aes-20170501-30-1002"), key, initialVector, padding...)
		if err != nil {
			t.Fatal(err)
		}
		s := base64.StdEncoding.EncodeToString(se)
		sed, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			t.Fatal(err)
		}
		sd, err := AESCBCDecrypt(sed, key, initialVector, padding...)
		if err != nil {
			t.Fatal(err)
		}
		if string(sd) != "aes-20170501-30-1002" {
			t.Log(string(sd))
			t.FailNow()
		}
	}
}

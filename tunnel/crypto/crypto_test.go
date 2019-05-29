package crypto

import (
	"testing"
)

func TestDes(t *testing.T) {
	des := NewCryptoDES("123")
	origData := []byte("flkajfklajfklsjafha;lfh")
	cryptedData, err := des.Encrypt(origData)
	if err != nil {
		t.Error(err)
		return
	}
	decryptedData, err := des.Decrypt(cryptedData)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(decryptedData))
}

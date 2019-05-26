package crypto

type CryptoBlank struct {}

func (c *CryptoBlank) Encrypt(b []byte) ([]byte, error) {
	return b, nil
}

func (c *CryptoBlank) Decrypt(b []byte) ([]byte, error) {
	return b, nil
}
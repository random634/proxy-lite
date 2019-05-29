package crypto

type CryptoMethod interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

type CryptoMethodType int32

const (
	CryptoMethodBlank CryptoMethodType = 1 + iota
	CryptoMethodDES
)

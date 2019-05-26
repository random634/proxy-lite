package tunnel

import (
	"github.com/random634/proxy-lite/src/tunnel/crypto"
)

type Tunnel struct {
	Addr     string
	Password string
	CryptoMethod *CryptoMethod
}

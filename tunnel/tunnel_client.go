package tunnel

import (
	"net"

	"github.com/random634/proxy-lite/tunnel/crypto"
)

type TunnelClient interface {
	Dial(network, addr string) (t Tunnel, err error)
	Close() (err error)
}

type TunnelClientImpl struct {
	CryptoMethod crypto.CryptoMethod
}

func NewTunnelClient(cryptoMethod crypto.CryptoMethod) TunnelClient {
	tc := new(TunnelClientImpl)
	tc.CryptoMethod = cryptoMethod

	return tc
}

func (tc *TunnelClientImpl) Dial(network, addr string) (t Tunnel, err error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	t = NewTunnel(conn, tc.CryptoMethod)

	return
}

func (ts *TunnelClientImpl) Close() (err error) {
	return nil
}

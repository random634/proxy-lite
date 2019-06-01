package tunnel

import (
	"net"

	"github.com/random634/proxy-lite/tunnel/crypto"
)

type TunnelServer interface {
	Listen(network, addr string) (err error)
	Accept() (t Tunnel, err error)
	Close() (err error)
}

type TunnelServerImpl struct {
	Listener     net.Listener
	CryptoMethod crypto.CryptoMethod
}

func NewTunnelServer(cryptoMethod crypto.CryptoMethod) TunnelServer {
	ts := new(TunnelServerImpl)
	ts.CryptoMethod = cryptoMethod

	return ts
}

func (ts *TunnelServerImpl) Listen(network, addr string) (err error) {
	listen, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	ts.Listener = listen

	return
}

func (ts *TunnelServerImpl) Accept() (t Tunnel, err error) {
	if ts.Listener == nil {
		return nil, ErrTunnelListenerNotExist
	}

	conn, err := ts.Listener.Accept()

	if err != nil {
		return nil, err
	}

	t = NewTunnel(conn, ts.CryptoMethod)

	return
}

func (ts *TunnelServerImpl) Close() (err error) {
	if ts.Listener == nil {
		return ts.Listener.Close()
	}

	return nil
}

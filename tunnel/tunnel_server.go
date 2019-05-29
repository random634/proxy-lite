package tunnel

import (
	"net"
)

func (t *Tunnel) Listen(network, addr string) error {
	if network != "" {
		t.Net = network
	}

	if addr != "" {
		t.Addr = addr
	}

	// host, strPort, err := net.SplitHostPort(t.Addr)
	// if err != nil {
	// 	return err
	// }
	// port, err := strconv.Atoi(strPort)
	// if err != nil {
	// 	return err
	// }

	listen, err := net.Listen(t.Net, t.Addr)
	if err != nil {
		return err
	}

	t.Listener = listen

	return err
}

func (t *Tunnel) Accept() (*Tunnel, error) {
	if t.Listener == nil {
		return nil, ErrTunnelListenerNotExist
	}

	conn, err := t.Listener.Accept()

	if err != nil {
		return nil, err
	}

	tunnel := &Tunnel{
		Net:          t.Net,
		Addr:         t.Addr,
		Password:     t.Password,
		CryptoMethod: t.CryptoMethod,
		Conn:         conn,
	}

	return tunnel, err
}

package tunnel

import (
	"net"
)

func (t *Tunnel) Dial(network, addr string) error {
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

	conn, err := net.Dial(t.Net, t.Addr)
	if err != nil {
		return err
	}

	t.Conn = conn

	return err
}

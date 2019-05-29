package tunnel

import (
	"errors"
	"io"
	"net"

	"github.com/random634/proxy-lite/tunnel/crypto"
)

type Tunnel struct {
	Net          string
	Addr         string
	Password     string
	CryptoMethod crypto.CryptoMethod
	Listener     net.Listener
	Conn         net.Conn
}

var (
	ErrTunnelNotExist         = errors.New("tunnel not exist")
	ErrTunnelListenerNotExist = errors.New("tunnel listener not exist")
	ErrTunnelConnNotExist     = errors.New("tunnel conn not exist")
)

func (t *Tunnel) Read() (b []byte, err error) {
	if t.Conn == nil {
		return nil, ErrTunnelConnNotExist
	}

	lenBuf := make([]byte, 2)
	io.ReadFull(t.Conn, lenBuf)

	// 小端序
	lenVal := int(lenBuf[0]) + int(lenBuf[1])*256
	dataBuf := make([]byte, lenVal)

	io.ReadFull(t.Conn, dataBuf)

	if t.CryptoMethod != nil {
		b, err = t.CryptoMethod.Decrypt(dataBuf)
		if err != nil {
			return nil, err
		}

		dataBuf = b
	}

	return dataBuf, nil
}

func (t *Tunnel) Write(dataBuf []byte) (n int, err error) {
	if t.Conn == nil {
		return 0, ErrTunnelConnNotExist
	}

	lenData := len(dataBuf)

	if t.CryptoMethod != nil {
		dataBuf, err = t.CryptoMethod.Encrypt(dataBuf)
		if err != nil {
			return 0, err
		}
	}

	lenDataEncrypt := len(dataBuf)

	lenBuf := make([]byte, 2)
	lenBuf[0] = byte(lenDataEncrypt % 256)
	lenBuf[1] = byte(lenDataEncrypt / 256)

	// byte array join
	buf := append(lenBuf, dataBuf...)

	n, err = t.Conn.Write(buf)

	return lenData, err
}

func (t *Tunnel) Close() error {
	if t.Listener != nil {
		t.Listener.Close()
		return nil
	}

	if t.Conn != nil {
		t.Conn.Close()
		return nil
	}

	return ErrTunnelNotExist
}

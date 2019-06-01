package tunnel

import (
	"testing"
	"time"

	"github.com/random634/proxy-lite/tunnel/crypto"
)

func TestTunnel(t *testing.T) {
	cryptoMethod := crypto.NewCryptoDES("tnt")

	server := func(ch chan int) {

		ts := NewTunnelServer(cryptoMethod)

		defer ts.Close()

		err := ts.Listen("tcp", "127.0.0.1:15050")
		if err != nil {
			t.Error(err)
			return
		}

		serveConn := func(tun Tunnel, ch chan int) {
			defer tun.Close()
			dataBuf, err := tun.Read()
			if err != nil {
				t.Error(err)
				return
			}

			t.Logf("server recv: %s", string(dataBuf))

			ch <- len(dataBuf)
		}

		// for {
		tun, err := ts.Accept()
		if err != nil {
			t.Log(err)
			return
		}

		chIn := make(chan int, 1)

		go serveConn(tun, chIn)
		// }

		chInside := <-chIn
		ch <- chInside
	}

	client := func() {
		tc := NewTunnelClient(cryptoMethod)

		defer tc.Close()

		tun, err := tc.Dial("tcp", "127.0.0.1:15050")
		if err != nil {
			t.Error(err)
			return
		}

		defer tun.Close()

		dataBuf := []byte("tnt is a kind of bomb, is it true?")
		n, err := tun.Write(dataBuf)
		if err != nil {
			t.Error(err)
			return
		}

		t.Logf("client write: %d bytes", n)
	}

	chOut := make(chan int, 1)

	go server(chOut)

	// let server run first
	time.Sleep(100)

	go client()

	outSide := <-chOut

	t.Logf("server recv: %d bytes", outSide)
}

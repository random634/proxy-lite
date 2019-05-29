package tunnel

import (
	"testing"
	"time"

	"github.com/random634/proxy-lite/tunnel/crypto"
)

func TestTunnel(t *testing.T) {
	server := func(ch chan int) {
		ts := &Tunnel{
			Net:          "tcp",
			Addr:         "127.0.0.1:15050",
			Password:     "tnt",
			CryptoMethod: crypto.NewCryptoDES("tnt"),
		}

		err := ts.Listen("", "")
		if err != nil {
			t.Error(err)
			return
		}

		serveConn := func(tun *Tunnel, ch chan int) {
			dataBuf, err := tun.Read()
			if err != nil {
				t.Error(err)
				return
			}

			t.Logf("server recv: %s", string(dataBuf))

			ch <- len(dataBuf)
		}

		// for {
		conn, err := ts.Accept()
		if err != nil {
			t.Log(err)
			return
		}

		chIn := make(chan int, 1)

		go serveConn(conn, chIn)
		// }

		chInside := <-chIn
		ch <- chInside
	}

	client := func() {
		tc := &Tunnel{
			Net:          "tcp",
			Addr:         "127.0.0.1:15050",
			Password:     "tnt",
			CryptoMethod: crypto.NewCryptoDES("tnt"),
		}

		err := tc.Dial("", "")
		if err != nil {
			t.Error(err)
			return
		}

		dataBuf := []byte("tnt is a kind of bomb, is it true?")
		n, err := tc.Write(dataBuf)
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

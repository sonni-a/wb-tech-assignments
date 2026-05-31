package client

import (
	"io"
	"net"
	"os"
	"sync"
	"time"
)

// Dial opens a TCP connection to host:port with the given dial timeout.
func Dial(host, port string, timeout time.Duration) (net.Conn, error) {
	dialer := net.Dialer{Timeout: timeout}
	return dialer.Dial("tcp", net.JoinHostPort(host, port))
}

// Relay copies bytes between STDIN and conn until EOF or the peer closes.
func Relay(conn net.Conn) {
	var closeOnce sync.Once
	closeConn := func() {
		closeOnce.Do(func() {
			_ = conn.Close()
		})
	}
	defer closeConn()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer closeConn()

		_, err := io.Copy(conn, os.Stdin)
		if err == nil {
			if tcp, ok := conn.(*net.TCPConn); ok {
				_ = tcp.CloseWrite()
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer closeConn()

		_, _ = io.Copy(os.Stdout, conn)
	}()

	wg.Wait()
}

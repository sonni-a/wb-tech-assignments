package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"telnet/internal/client"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: telnet [--timeout duration] <host> <port>")
		os.Exit(2)
	}

	host, port := args[0], args[1]

	conn, err := client.Dial(host, port, *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connection failed: %v\n", err)
		os.Exit(1)
	}

	client.Relay(conn)
}

package probes

import (
	"net"
	"testing"
	"time"
)

func TestTCPProbe(t *testing.T) {
	c := ProbeConfig{
		Target:  "/tmp/go_test.sock",
		Timeout: time.Duration(1 * time.Second),
	}

	var nw string
	var add string
	var to time.Duration

	f := func(network string, address string, timeout time.Duration) (net.Conn, error) {
		nw = network
		add = address
		to = timeout

		conn, _ := net.Dial("unix", address)

		return conn, nil
	}

	tcp := TCPProbe{
		config:      c,
		DialTimeout: f,
	}

	tcp.Execute()
}

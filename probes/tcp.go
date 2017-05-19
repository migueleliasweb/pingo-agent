package probes

import (
	"net"
	"time"
)

//TCPProbe Struct for probing TCP
type TCPProbe struct {
	DialTimeout func(string, string, time.Duration) (net.Conn, error)
	config      ProbeConfig
}

//Execute Executes TCP probing
func (probe *TCPProbe) Execute() (uint8, error) {
	startTime := time.Now()
	conn, err := probe.DialTimeout(
		"tcp",
		probe.config.Target,
		probe.config.Timeout,
	)

	if err != nil {
		return uint8(0), err
	}

	duration := uint8(time.Since(startTime) / time.Nanosecond)
	conn.Close()

	return duration, err
}

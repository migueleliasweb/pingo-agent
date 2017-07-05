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
func (probe TCPProbe) Execute() (time.Duration, error) {
	startTime := time.Now()
	conn, err := probe.DialTimeout(
		"tcp",
		probe.config.Target,
		probe.config.Timeout,
	)

	defer func() {
		// not sure why this is needed
		//apparentely just some types of connections
		//implement the Close method
		switch conn.(type) {
		case net.Conn:
			conn.Close()
		}
	}()

	if err != nil {
		return time.Duration(0), err
	}

	duration := time.Duration(time.Since(startTime) / time.Nanosecond)

	return duration, err
}

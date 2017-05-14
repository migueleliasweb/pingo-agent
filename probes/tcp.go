package probes

import (
	"net"
	"time"
)

func TcpCheck(config ProbeConfig) (uint8, error) {
	startTime := time.Now()
	conn, err := net.DialTimeout(
		"tcp",
		config.GetTarget(),
		config.GetTimeout(),
	)

	if err != nil {
		return uint8(0), err
	}

	duration := uint8(time.Now().Sub(startTime) / time.Nanosecond)
	conn.Close()

	return duration, err
}

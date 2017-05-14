package probes

import (
	"net"
	"time"
)

func DnsProbe(config ProbeConfig) (uint8, error, []string) {
	startTime := time.Now()
	addrs, err := net.LookupHost(config.GetTarget())
	duration := uint8(time.Now().Sub(startTime) / time.Nanosecond)

	return duration, err, addrs
}

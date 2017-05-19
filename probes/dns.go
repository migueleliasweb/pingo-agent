package probes

import (
	"net"
	"time"
)

//DNSProbe Probe for DNS
type DNSProbe struct {
	config     ProbeConfig
	LookupHost func(string) ([]string, error)
}

//Execute DNS probing
func (probe *DNSProbe) Execute() (uint8, error) {
	startTime := time.Now()
	_, err := probe.LookupHost(probe.config.Target)
	duration := uint8(time.Now().Sub(startTime) / time.Nanosecond)

	return duration, err
}

//NewDNSProbe Returns new DNSProbe instance
func NewDNSProbe(config ProbeConfig) *DNSProbe {
	dnsProbe := DNSProbe{
		config:     config,
		LookupHost: net.LookupHost,
	}

	return &dnsProbe
}

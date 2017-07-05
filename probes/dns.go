package probes

import (
	"net"
	"time"
)

//DNSProbe Probe for DNS
type DNSProbe struct {
	config         ProbeConfig
	lookupHostFunc func(string) ([]string, error)
}

//Execute DNS probing
func (probe *DNSProbe) Execute() (time.Duration, error) {
	startTime := time.Now()
	_, err := probe.lookupHostFunc(probe.config.Target)
	duration := time.Duration(time.Now().Sub(startTime) / time.Nanosecond)

	return duration, err
}

//NewDNSProbe Returns new DNSProbe instance
func NewDNSProbe(config ProbeConfig) *DNSProbe {
	dnsProbe := DNSProbe{
		config:         config,
		lookupHostFunc: net.LookupHost,
	}

	return &dnsProbe
}

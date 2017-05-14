package probes

import "time"

type IProbeConfig interface {
	GetTarget() string
	GetTimeout() time.Duration
}

type ProbeConfig struct {
	Target  string
	Timeout time.Duration
}

func (c *ProbeConfig) GetTarget() string {
	return c.Target
}

func (c *ProbeConfig) GetTimeout() time.Duration {
	return time.Duration(c.Timeout) * time.Second
}

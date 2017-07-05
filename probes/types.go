package probes

import "time"

//Probe Interface for all Probes
type Probe interface {
	Execute() (time.Duration, error)
}

//ProbeConfig Configuration for all Probes
type ProbeConfig struct {
	Target  string
	Timeout time.Duration
}

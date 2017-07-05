package probes

import (
	"reflect"
	"testing"
	"time"
)

var called bool

func fakeLookupHost(string) ([]string, error) {
	called = true
	return []string{}, nil
}

func TestDNSProbe(t *testing.T) {
	config := ProbeConfig{
		Target:  "localhost:1234",
		Timeout: time.Duration(1 * time.Second),
	}

	p := DNSProbe{
		config:         config,
		lookupHostFunc: fakeLookupHost,
	}

	p.Execute()

	if called != true {
		t.Error("The lookupHostFunc wasn't called.")
	}
}

func TestNewDNSProbe(t *testing.T) {
	c := ProbeConfig{
		Target:  "localhost:1234",
		Timeout: time.Duration(1 * time.Second),
	}

	p := NewDNSProbe(c)

	if !reflect.DeepEqual(c, p.config) {
		t.Error("Wrong configuration was set on NewDNSProbe")
	}

	//Can't test properly this assertion because the net.LookupHost is
	//plattform specific and relies on local implementations
	// if !reflect.DeepEqual(net.LookupHost, p.lookupHostFunc) {
	// 	t.Error("Wrong lookupHostFunc function")
	// }
}

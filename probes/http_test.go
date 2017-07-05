package probes

import (
	"net/http"
	"testing"
	"time"
)

var target string
var calledGet bool

type fakeHTTPClient struct {
	target *string
}

func (client fakeHTTPClient) Get(s string) (*http.Response, error) {
	target = s
	calledGet = true

	return &http.Response{}, nil
}

func TestHTTPProbe(t *testing.T) {
	p := HTTPProbe{
		config: ProbeConfig{
			Target:  "localhost:1234",
			Timeout: time.Duration(1 * time.Second),
		},
		httpClient: fakeHTTPClient{},
	}

	p.Execute()

	if !calledGet {
		t.Error("Get method not called")
	}

	if target != p.config.Target {
		t.Error("Wrong target on HTTPProbe")
	}
}

func TestNewHTTPProbe(t *testing.T) {
	c := ProbeConfig{
		Target:  "localhost:1234",
		Timeout: time.Duration(1 * time.Second),
	}

	p := NewHTTPProbe(c)

	if p.config.Target != c.Target {
		t.Error("Wrong target when using NewHTTPProbe")
	}

	if p.config.Timeout != c.Timeout {
		t.Error("Wrong timeout when using NewHTTPProbe")
	}
}

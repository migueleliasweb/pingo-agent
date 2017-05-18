package probes

import (
	"net/http"
	"time"
)

//Client Interface to help testing http requests
type Client interface {
	Get(string) (*http.Response, error)
}

//HttpProbe Struct for the HttpProbe
type HttpProbe struct {
	config     ProbeConfig
	httpClient Client
}

func (probe *HttpProbe) Execute() (uint8, error) {
	startTime := time.Now()
	_, err := probe.httpClient.Get(probe.config.Target)
	duration := uint8(time.Now().Sub(startTime) / time.Nanosecond)

	return duration, err
}

func New(config ProbeConfig) *HttpProbe {
	// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	var client = &http.Client{Timeout: config.Timeout}

	httpProbe := HttpProbe{
		httpClient: client,
		config:     config,
	}

	return &httpProbe
}

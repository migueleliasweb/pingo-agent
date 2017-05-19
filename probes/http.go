package probes

import (
	"net/http"
	"time"
)

//Client Interface to help testing http requests
type Client interface {
	Get(string) (*http.Response, error)
}

//HTTPProbe Struct for the HTTPProbe
type HTTPProbe struct {
	config     ProbeConfig
	httpClient Client
}

//Execute Executes HTTPProbe
func (probe *HTTPProbe) Execute() (uint8, error) {
	startTime := time.Now()
	_, err := probe.httpClient.Get(probe.config.Target)
	duration := uint8(time.Now().Sub(startTime) / time.Nanosecond)

	return duration, err
}

//NewHTTPProbe Returns a new instance of HTTPProbe
func NewHTTPProbe(config ProbeConfig) *HTTPProbe {
	// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	var client = &http.Client{Timeout: config.Timeout}

	httpProbe := HTTPProbe{
		httpClient: client,
		config:     config,
	}

	return &httpProbe
}

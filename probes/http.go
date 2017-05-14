package probes

import (
	"net/http"
	"time"
)

func HttpProbe(config ProbeConfig) (uint8, error, *http.Response) {
	// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	var netClient = &http.Client{Timeout: config.GetTimeout()}

	startTime := time.Now()
	response, err := netClient.Get(config.GetTarget())
	duration := uint8(time.Now().Sub(startTime) / time.Nanosecond)

	return duration, err, response
}

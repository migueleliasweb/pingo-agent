package main

import (
	"fmt"

	"github.com/migueleliasweb/pingo-agent/probes"
)

func main() {
	config := probes.ProbeConfig{
		Target:  "httpbin.org",
		Timeout: 2000,
	}

	duration, err, ips := probes.DnsProbe(config)

	fmt.Println("Error: ", err)
	fmt.Println("Duration: ", duration)
	fmt.Println("IPs: ", ips)

	config2 := probes.ProbeConfig{
		Target:  "httpbin.org:443",
		Timeout: 2000,
	}

	duration, err = probes.TcpCheck(config2)

	fmt.Println("Error: ", err)
	fmt.Println("Duration: ", duration)
}

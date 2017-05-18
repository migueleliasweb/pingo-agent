package main

import "github.com/migueleliasweb/pingo-agent/ws"
import "github.com/migueleliasweb/pingo-agent/probes"

func main() {
	domain := "sockb.in"
	path := "/repeat/5"
	tags := []string{}
	tags[0] = "foo"
	tags[1] = "bar"

    handlerMap := make(map[string]string)
    handlerMap["dns"] = probes.DnsProbe
    handlerMap["http"] = probes.HttpProbe
    handlerMap["tcp"] = probes.TcpCheck

	wsClient := ws.New()
	wsClient.Setup(
		domain,
		path,
		tags,
        handlerMap
	)
}

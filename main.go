package main

import "github.com/migueleliasweb/pingo-agent/ws"

func main() {
	domain := "sockb.in"
	path := "/repeat/5"
	tags := []string{}
	tags[0] = "foo"
	tags[1] = "bar"

	wsClient := ws.New()
	wsClient.Setup(
		domain,
		path,
		tags,
	)
}

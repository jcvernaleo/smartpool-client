package main

import (
	"encoding/json"
	"fmt"
)

type jsonRequest struct {
	Method  string          `json:"method"`
	Version string          `json:"jsonrpc"`
	Id      json.RawMessage `json:"id,omitempty"`
	Payload json.RawMessage `json:"params,omitempty"`
}

const (
	farmUrl = "localhost:1633"
	id      = "/:testminer/"
)

func main() {
	fmt.Println("Connecting to", farmUrl+id)

	m := "eth_submitWork"
	v := "2.0"
	i := json.RawMessage("1")
	p := json.RawMessage(`["0x8c4b491d264e81ca", "0xbdaf505164d6e215b76f87a1561be0ef7242c0bdef79f2377c942a0fd57fd110", "0xdfef46fd1600da329d7ac2c9b0e1c2aa046043a11c7a3994169bd39c5a41b729"]`)

	fakeSubmit := jsonRequest{
		Method:  m,
		Version: v,
		Id:      i,
		Payload: p,
	}

	fmt.Println(fakeSubmit.Method, fakeSubmit.Version, string(fakeSubmit.Id), string(fakeSubmit.Payload))

	//eth_submitHashrate 2.0 1 ["0x0","0xd1704618c93ff275e81ce2e224ca50f79199214c64cb06044f8f764409ef5db3"]
	//eth_getWork 2.0 1

	return
}

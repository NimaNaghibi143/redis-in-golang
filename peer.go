package main

import "net"

type Peer struct {
	conn net.Conn
}

func (p *Peer) readLoop() {
	for {

	}
}

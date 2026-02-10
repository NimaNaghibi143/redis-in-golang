package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
)

// port configuration
const defaultListenAddr = ":5001"

// defines how the server should run
type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	peers     map[*Peer]bool
	ln        net.Listener
	addPeerCh chan *Peer
}

// factory method, enduring the server has a valid port to listen on
func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListenAddr
	}
	return &Server{
		Config:    cfg,
		peers:     make(map[*Peer]bool),
		addPeerCh: make(chan *Peer),
	}
}

// Opens the socket and resolves the port on the OS
func (s *Server) start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln

	go s.loop()

	return s.acceptLoop()
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.addPeerCh:
			s.peers[peer] = true
		default:
			fmt.Println("Yo")
		}
	}
}

// An infinit loop for accepting a connection and using GOROUTINE to keep the conn in the background
func (s *Server) acceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("accept error", "err", err)
			continue
		}
		go s.handleConn(conn)
	}
}

// we are going to handle the conn, we make a new peer and add this peer to the peer channel for maintenance
func (s *Server) handleConn(conn net.Conn) {
	peer := NewPeer(conn)
	s.addPeerCh <- peer
	peer.readLoop()
}

func main() {
	server := NewServer(Config{})
	log.Fatal(server.start())
}

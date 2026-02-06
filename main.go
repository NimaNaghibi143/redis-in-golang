package main

import (
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
	ln net.Listener
}

// factory method, enduring the server has a valid port to listen on
func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListenAddr
	}
	return &Server{
		Config: cfg,
	}
}

func (s *Server) start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln

	return s.acceptLoop()
}

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

func (s *Server) handleConn(conn net.Conn) {

}

func main() {

}

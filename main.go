package main

import (
	"fmt"

	"log"
	"log/slog"
	"net"
	"strconv"
)

const defaultListenAddr = ":5001"

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	peers     map[*Peer]bool
	Ln        net.Listener
	addPeerCh chan *Peer
	quitCh    chan struct{}
	msgCh     chan []byte
}

func NewServer(cfg Config) *Server {
	if cfg.ListenAddr == "" {
		cfg.ListenAddr = defaultListenAddr
	}

	return &Server{
		Config:    cfg,
		peers:     make(map[*Peer]bool),
		addPeerCh: make(chan *Peer),
		quitCh:    make(chan struct{}),
		msgCh:     make(chan []byte),
	}
}

func (s *Server) Start() error {
	Ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.Ln = Ln

	go s.loop()

	slog.Info("server running", "listenAddr", s.ListenAddr)

	return s.acceptLoop()
}

func (s Server) loop() {
	for {
		select {
		case rawMsg := <-s.msgCh:
			if err := s.handleRawMessage(rawMsg); err != nil {
				slog.Error("handle Raw Message error", "err", err)
			}
		case <-s.quitCh:
			return
		case peer := <-s.addPeerCh:
			s.peers[peer] = true
		}
	}
}

func (s *Server) acceptLoop() error {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			slog.Error("accept error", "err", err)
			continue
		}
		go s.handleConn(conn)
	}

}

func (s *Server) handleConn(conn net.Conn) {
	peer := NewPeer(
		conn,
		s.msgCh,
	)
	s.addPeerCh <- peer
	slog.Info("new peer connected", "remoteAddr", conn.RemoteAddr())
	if err := peer.readLoop(); err != nil {
		slog.Error("peer read error", "err", err, "remoteAddr", conn.RemoteAddr())
	}
}

func (s *Server) handleRawMessage(msg []byte) error {

	// get err
	strMsg, err := strconv.Unquote(string(msg))
	if err != nil {
		return fmt.Errorf("unquote raw msg error: %w", err)
	}
	cmd, err := ParseCommand(strMsg)
	if err != nil {
		return err
	}
	slog.Info("received command", "cmd", cmd)
	return nil
}

func main() {

	go func() {
		server := NewServer(Config{ListenAddr: defaultListenAddr})
		log.Fatal(server.Start())
	}()

	select {} // blocking here
}

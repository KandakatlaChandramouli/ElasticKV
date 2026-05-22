package transport

import (
	"io"
	"net"
	"sync/atomic"
)

type Server struct {
	Address  string
	Received atomic.Uint64
}

func NewServer(
	address string,
) *Server {

	return &Server{
		Address: address,
	}
}

func (s *Server) Start() error {

	listener, err := net.Listen(
		"tcp",
		s.Address,
	)

	if err != nil {
		return err
	}

	go func() {

		for {

			conn, err := listener.Accept()

			if err != nil {
				continue
			}

			go s.handleConn(conn)
		}
	}()

	return nil
}

func (s *Server) handleConn(
	conn net.Conn,
) {

	defer conn.Close()

	buffer := make([]byte, 8192)

	for {

		n, err := conn.Read(buffer)

		if err != nil {

			if err == io.EOF {
				return
			}

			return
		}

		_, err = DecodeFrame(
			buffer[:n],
		)

		if err != nil {
			continue
		}

		s.Received.Add(1)
	}
}

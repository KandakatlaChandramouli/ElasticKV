package transport

import (
	"bufio"
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

	reader := bufio.NewReaderSize(
		conn,
		64*1024,
	)

	for {

		_, err := DecodeFrame(reader)

		if err != nil {

			if err == io.EOF {
				return
			}

			return
		}

		s.Received.Add(1)
	}
}

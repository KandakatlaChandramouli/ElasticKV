package transport

import (
	"bufio"
	"net"
)

type Client struct {
	Connection net.Conn
	Writer     *bufio.Writer
}

func NewClient(
	address string,
) (*Client, error) {

	conn, err := net.Dial(
		"tcp",
		address,
	)

	if err != nil {
		return nil, err
	}

	return &Client{
		Connection: conn,
		Writer: bufio.NewWriterSize(
			conn,
			64*1024,
		),
	}, nil
}

func (c *Client) Send(
	frame ReplicationFrame,
) error {

	data, err := EncodeFrame(
		frame,
	)

	if err != nil {
		return err
	}

	_, err = c.Writer.Write(data)

	if err != nil {
		return err
	}

	return c.Writer.Flush()
}

func (c *Client) Close() error {
	return c.Connection.Close()
}

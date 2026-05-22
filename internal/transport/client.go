package transport

import (
	"net"
)

type Client struct {
	Connection net.Conn
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

	_, err = c.Connection.Write(data)

	return err
}

func (c *Client) Close() error {
	return c.Connection.Close()
}

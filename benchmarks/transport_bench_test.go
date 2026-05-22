package benchmarks

import (
	"bytes"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/transport"
)

func getFreePort() (string, error) {

	listener, err := net.Listen(
		"tcp",
		":0",
	)

	if err != nil {
		return "", err
	}

	defer listener.Close()

	return fmt.Sprintf(
		":%d",
		listener.Addr().(*net.TCPAddr).Port,
	), nil
}

func BenchmarkTransportReplication(
	b *testing.B,
) {

	address, err := getFreePort()

	if err != nil {
		b.Fatal(err)
	}

	server := transport.NewServer(
		address,
	)

	err = server.Start()

	if err != nil {
		b.Fatal(err)
	}

	time.Sleep(100 * time.Millisecond)

	client, err := transport.NewClient(
		"127.0.0.1" + address,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer client.Close()

	payload := bytes.Repeat(
		[]byte("R"),
		1024,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		frame := transport.ReplicationFrame{
			SequenceID: uint64(i),
			Payload:    payload,
		}

		err := client.Send(frame)

		if err != nil {
			b.Fatal(err)
		}
	}

	time.Sleep(200 * time.Millisecond)

	b.StopTimer()

	b.Logf(
		"Frames Received = %d",
		server.Received.Load(),
	)
}

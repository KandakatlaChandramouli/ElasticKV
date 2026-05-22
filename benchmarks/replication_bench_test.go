package benchmarks

import (
	"bytes"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/transport"
)

func getReplicationPort() (string, error) {

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

func BenchmarkAsyncReplication(
	b *testing.B,
) {

	address, err := getReplicationPort()

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

	replicator := transport.NewReplicator(
		client,
		8192,
	)

	replicator.Start()

	defer replicator.Stop()

	payload := bytes.Repeat(
		[]byte("A"),
		1024,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		frame := transport.ReplicationFrame{
			SequenceID: uint64(i),
			Payload:    payload,
		}

		replicator.Replicate(frame)
	}

	time.Sleep(500 * time.Millisecond)

	b.StopTimer()

	b.Logf(
		"Replicated = %d",
		replicator.Sent.Load(),
	)

	b.Logf(
		"Dropped = %d",
		replicator.Dropped.Load(),
	)

	b.Logf(
		"Server Received = %d",
		server.Received.Load(),
	)
}

package benchmarks

import (
	"bytes"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/transport"
)

func getWireBatchPort() (string, error) {

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

func BenchmarkWireBatchReplication(
	b *testing.B,
) {

	address, err := getWireBatchPort()

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

	replicator := transport.NewWireBatchReplicator(
		client,
		131072,
		256,
		5*time.Millisecond,
	)

	replicator.Start()

	defer replicator.Stop()

	payload := bytes.Repeat(
		[]byte("W"),
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

	b.StopTimer()

	time.Sleep(2 * time.Second)

	b.Logf(
		"Sent = %d",
		replicator.Sent.Load(),
	)

	b.Logf(
		"Dropped = %d",
		replicator.Dropped.Load(),
	)

	b.Logf(
		"Batches = %d",
		replicator.Batches.Load(),
	)

	b.Logf(
		"Server Received = %d",
		server.Received.Load(),
	)
}

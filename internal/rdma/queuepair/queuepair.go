package queuepair

type QueuePair struct {
	SendDepth    int
	ReceiveDepth int
}

func New(
	send int,
	receive int,
) QueuePair {

	return QueuePair{
		SendDepth:    send,
		ReceiveDepth: receive,
	}
}

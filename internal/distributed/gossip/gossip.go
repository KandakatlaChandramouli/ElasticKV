package gossip

type Message struct {
	Node  string
	State string
}

func Broadcast(
	nodes []string,
	state string,
) []Message {

	messages := make([]Message, 0)

	for _, node := range nodes {

		messages = append(
			messages,
			Message{
				Node:  node,
				State: state,
			},
		)
	}

	return messages
}

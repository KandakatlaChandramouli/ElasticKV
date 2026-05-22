package requestqueue

type Request struct {
	ID      int
	Payload []byte
}

type Queue struct {
	requests []Request
}

func NewQueue() *Queue {
	return &Queue{
		requests: make([]Request, 0),
	}
}

func (q *Queue) Push(
	request Request,
) {
	q.requests = append(
		q.requests,
		request,
	)
}

func (q *Queue) Count() int {
	return len(q.requests)
}

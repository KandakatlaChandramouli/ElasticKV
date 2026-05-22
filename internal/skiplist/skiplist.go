package skiplist

import "sync"

type Node struct {
	Key   uint64
	Value []byte
	Next  *Node
}

type List struct {
	Head  *Node
	Mutex sync.RWMutex
}

func New() *List {

	return &List{}
}

func (l *List) Insert(
	key uint64,
	value []byte,
) {

	l.Mutex.Lock()

	defer l.Mutex.Unlock()

	node := &Node{
		Key:   key,
		Value: value,
		Next:  l.Head,
	}

	l.Head = node
}

func (l *List) Lookup(
	key uint64,
) ([]byte, bool) {

	l.Mutex.RLock()

	defer l.Mutex.RUnlock()

	current := l.Head

	for current != nil {

		if current.Key == key {
			return current.Value, true
		}

		current = current.Next
	}

	return nil, false
}

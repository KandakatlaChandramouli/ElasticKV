package cache

import (
	"container/list"
)

type Entry struct {
	Key   uint64
	Value []byte
}

type LRU struct {
	Capacity int
	List     *list.List
	Items    map[uint64]*list.Element
}

func NewLRU(
	capacity int,
) *LRU {

	return &LRU{
		Capacity: capacity,
		List:     list.New(),
		Items:    make(map[uint64]*list.Element),
	}
}

func (l *LRU) Put(
	key uint64,
	value []byte,
) {

	if element, ok := l.Items[key]; ok {

		l.List.MoveToFront(element)

		element.Value = Entry{
			Key:   key,
			Value: value,
		}

		return
	}

	element := l.List.PushFront(
		Entry{
			Key:   key,
			Value: value,
		},
	)

	l.Items[key] = element

	if l.List.Len() > l.Capacity {

		tail := l.List.Back()

		if tail == nil {
			return
		}

		l.List.Remove(tail)

		entry := tail.Value.(Entry)

		delete(
			l.Items,
			entry.Key,
		)
	}
}

func (l *LRU) Get(
	key uint64,
) ([]byte, bool) {

	element, ok := l.Items[key]

	if !ok {
		return nil, false
	}

	l.List.MoveToFront(element)

	entry := element.Value.(Entry)

	return entry.Value, true
}

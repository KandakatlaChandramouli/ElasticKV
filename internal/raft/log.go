package raft

import "sync"

type Entry struct {
	Index uint64
	Term uint64
	Data []byte
}

type Log struct {
	Entries []Entry
	Mutex sync.RWMutex
}

func NewLog() *Log {

	return &Log{
		Entries: make([]Entry, 0),
	}
}

func (l *Log) Append(
	entry Entry,
) {

	l.Mutex.Lock()

	defer l.Mutex.Unlock()

	l.Entries = append(
		l.Entries,
		entry,
	)
}

func (l *Log) LastIndex() uint64 {

	l.Mutex.RLock()

	defer l.Mutex.RUnlock()

	return uint64(len(l.Entries))
}

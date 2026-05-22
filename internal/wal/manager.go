package wal

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type Manager struct {
	Directory      string
	MaxSegmentSize uint64
	ActiveID       uint64
	ActiveSegment  *Segment
	Index          *SparseIndex
	Mutex          sync.Mutex
}

func OpenManager(
	directory string,
	maxSegmentSize uint64,
) (*Manager, error) {

	err := os.MkdirAll(
		directory,
		0755,
	)

	if err != nil {
		return nil, err
	}

	manager := &Manager{
		Directory:      directory,
		MaxSegmentSize: maxSegmentSize,
		Index: NewSparseIndex(
			1024,
		),
	}

	err = manager.rotate()

	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) segmentPath(
	id uint64,
) string {

	return filepath.Join(
		m.Directory,
		fmt.Sprintf(
			"%020d.wal",
			id,
		),
	)
}

func (m *Manager) rotate() error {

	if m.ActiveSegment != nil {

		err := m.ActiveSegment.Sync()

		if err != nil {
			return err
		}

		err = m.ActiveSegment.Close()

		if err != nil {
			return err
		}
	}

	path := m.segmentPath(
		m.ActiveID,
	)

	segment, err := OpenSegment(
		path,
	)

	if err != nil {
		return err
	}

	m.ActiveSegment = segment

	m.ActiveID++

	return nil
}

func (m *Manager) Append(
	entry Entry,
) error {

	m.Mutex.Lock()

	defer m.Mutex.Unlock()

	estimatedSize := uint64(
		HeaderSize + len(entry.Payload),
	)

	if m.ActiveSegment.Offset+
		estimatedSize >= m.MaxSegmentSize {

		err := m.rotate()

		if err != nil {
			return err
		}
	}

	offset, err := m.ActiveSegment.Append(
		entry,
	)

	if err != nil {
		return err
	}

	m.Index.Add(
		entry.SequenceID,
		offset,
		m.ActiveID-1,
	)

	return nil
}

func (m *Manager) Sync() error {

	m.Mutex.Lock()

	defer m.Mutex.Unlock()

	return m.ActiveSegment.Sync()
}

func (m *Manager) Close() error {

	m.Mutex.Lock()

	defer m.Mutex.Unlock()

	if m.ActiveSegment == nil {
		return nil
	}

	return m.ActiveSegment.Close()
}

func (m *Manager) Segments() ([]string, error) {

	files, err := filepath.Glob(
		filepath.Join(
			m.Directory,
			"*.wal",
		),
	)

	if err != nil {
		return nil, err
	}

	sort.Strings(files)

	return files, nil
}

func (m *Manager) Replay(
	handler func(Entry) error,
) error {

	segments, err := m.Segments()

	if err != nil {
		return err
	}

	for _, segment := range segments {

		err := Replay(
			segment,
			handler,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) ReplayFrom(
	sequenceID uint64,
	handler func(Entry) error,
) error {

	entry, ok := m.Find(
		sequenceID,
	)

	if !ok {

		return m.Replay(handler)
	}

	path := m.segmentPath(
		entry.SegmentID,
	)

	return ReplayFromOffset(
		path,
		entry.Offset,
		handler,
	)
}

func (m *Manager) Find(
	sequenceID uint64,
) (IndexEntry, bool) {

	return m.Index.Find(
		sequenceID,
	)
}

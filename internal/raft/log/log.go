package log

type Entry struct {
	Term    int
	Index   int
	Command []byte
}

type RaftLog struct {
	entries []Entry
}

func New() *RaftLog {
	return &RaftLog{
		entries: make([]Entry, 0),
	}
}

func (r *RaftLog) Append(
	entry Entry,
) {

	r.entries = append(
		r.entries,
		entry,
	)
}

func (r *RaftLog) Count() int {
	return len(r.entries)
}

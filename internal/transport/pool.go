package transport

import (
	"bytes"
	"sync"
)

var BufferPool = sync.Pool{
	New: func() any {

		buffer := bytes.NewBuffer(
			make([]byte, 0, 1024*1024),
		)

		return buffer
	},
}

func AcquireBuffer() *bytes.Buffer {

	buffer := BufferPool.Get().(*bytes.Buffer)

	buffer.Reset()

	return buffer
}

func ReleaseBuffer(
	buffer *bytes.Buffer,
) {

	if buffer == nil {
		return
	}

	buffer.Reset()

	BufferPool.Put(buffer)
}

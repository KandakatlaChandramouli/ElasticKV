package pagecache

import "sync"

type Page struct {
	ID   uint64
	Data []byte
}

type Cache struct {
	Pages map[uint64]Page
	Mutex sync.RWMutex
}

func New() *Cache {

	return &Cache{
		Pages: make(
			map[uint64]Page,
		),
	}
}

func (c *Cache) Put(
	id uint64,
	data []byte,
) {

	c.Mutex.Lock()

	defer c.Mutex.Unlock()

	c.Pages[id] = Page{
		ID:   id,
		Data: data,
	}
}

func (c *Cache) Get(
	id uint64,
) ([]byte, bool) {

	c.Mutex.RLock()

	defer c.Mutex.RUnlock()

	page, ok := c.Pages[id]

	if !ok {
		return nil, false
	}

	return page.Data, true
}

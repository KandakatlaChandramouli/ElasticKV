package pagecache

type Cache struct {
	pages map[int][]byte
}

func NewCache() *Cache {
	return &Cache{
		pages: make(map[int][]byte),
	}
}

func (c *Cache) Put(
	id int,
	data []byte,
) {
	c.pages[id] = data
}

func (c *Cache) Get(
	id int,
) ([]byte, bool) {

	value, ok := c.pages[id]

	return value, ok
}

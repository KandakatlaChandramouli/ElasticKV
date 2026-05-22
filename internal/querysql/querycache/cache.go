package querycache

type Cache struct {
	queries map[string]string
}

func NewCache() *Cache {
	return &Cache{
		queries: make(map[string]string),
	}
}

func (c *Cache) Put(
	query string,
	result string,
) {
	c.queries[query] = result
}

func (c *Cache) Get(
	query string,
) (string, bool) {

	value, ok := c.queries[query]

	return value, ok
}

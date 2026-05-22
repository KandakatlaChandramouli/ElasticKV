package querycache

type Cache struct {
	values map[string][]int
}

func NewCache() *Cache {
	return &Cache{
		values: make(map[string][]int),
	}
}

func (c *Cache) Put(
	key string,
	ids []int,
) {
	c.values[key] = ids
}

func (c *Cache) Get(
	key string,
) ([]int, bool) {

	value, ok := c.values[key]

	return value, ok
}

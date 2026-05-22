package cache

type Cache struct {
	vectors map[string][]float32
}

func NewCache() *Cache {
	return &Cache{
		vectors: make(map[string][]float32),
	}
}

func (c *Cache) Put(
	key string,
	vector []float32,
) {
	c.vectors[key] = vector
}

func (c *Cache) Get(
	key string,
) ([]float32, bool) {

	value, ok := c.vectors[key]

	return value, ok
}

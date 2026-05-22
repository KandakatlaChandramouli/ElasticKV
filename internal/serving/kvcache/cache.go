package kvcache

type KVCache struct {
	values map[string][]float32
}

func NewCache() *KVCache {
	return &KVCache{
		values: make(map[string][]float32),
	}
}

func (c *KVCache) Put(
	key string,
	value []float32,
) {
	c.values[key] = value
}

func (c *KVCache) Get(
	key string,
) ([]float32, bool) {

	value, ok := c.values[key]

	return value, ok
}

package cache

import "sync"

type Sharded struct {
	Shards []*Shard
	Count  uint64
}

type Shard struct {
	Cache *LRU
	Mutex sync.Mutex
}

func NewSharded(
	shards uint64,
	capacity int,
) *Sharded {

	result := make(
		[]*Shard,
		0,
		shards,
	)

	perShard := capacity / int(shards)

	for i := uint64(0); i < shards; i++ {

		result = append(
			result,
			&Shard{
				Cache: NewLRU(perShard),
			},
		)
	}

	return &Sharded{
		Shards: result,
		Count:  shards,
	}
}

func (s *Sharded) shard(
	key uint64,
) *Shard {

	return s.Shards[key%s.Count]
}

func (s *Sharded) Put(
	key uint64,
	value []byte,
) {

	shard := s.shard(key)

	shard.Mutex.Lock()

	defer shard.Mutex.Unlock()

	shard.Cache.Put(
		key,
		value,
	)
}

func (s *Sharded) Get(
	key uint64,
) ([]byte, bool) {

	shard := s.shard(key)

	shard.Mutex.Lock()

	defer shard.Mutex.Unlock()

	return shard.Cache.Get(key)
}

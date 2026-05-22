package shard

func Route(
	key uint64,
	shardCount uint64,
) uint64 {
	return key % shardCount
}

package merkle

import "hash/fnv"

func Hash(
	data []byte,
) uint64 {

	hash := fnv.New64a()

	hash.Write(data)

	return hash.Sum64()
}

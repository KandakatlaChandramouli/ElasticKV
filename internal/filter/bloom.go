package filter

import (
	"hash/fnv"
)

type BloomFilter struct {
	Bits []bool
	Size uint64
}

func NewBloomFilter(
	size uint64,
) *BloomFilter {

	return &BloomFilter{
		Bits: make(
			[]bool,
			size,
		),
		Size: size,
	}
}

func hash1(
	key uint64,
) uint64 {

	h := fnv.New64a()

	var b [8]byte

	for i := 0; i < 8; i++ {
		b[i] = byte(
			key >> (i * 8),
		)
	}

	h.Write(b[:])

	return h.Sum64()
}

func hash2(
	key uint64,
) uint64 {

	x := key

	x ^= x >> 33
	x *= 0xff51afd7ed558ccd
	x ^= x >> 33
	x *= 0xc4ceb9fe1a85ec53
	x ^= x >> 33

	return x
}

func (b *BloomFilter) Add(
	key uint64,
) {

	h1 := hash1(key) % b.Size
	h2 := hash2(key) % b.Size

	b.Bits[h1] = true
	b.Bits[h2] = true
}

func (b *BloomFilter) MayContain(
	key uint64,
) bool {

	h1 := hash1(key) % b.Size
	h2 := hash2(key) % b.Size

	return b.Bits[h1] &&
		b.Bits[h2]
}

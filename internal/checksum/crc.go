package checksum

import "hash/crc32"

func Compute(
	data []byte,
) uint32 {

	return crc32.ChecksumIEEE(
		data,
	)
}

func Verify(
	data []byte,
	sum uint32,
) bool {

	return Compute(data) == sum
}

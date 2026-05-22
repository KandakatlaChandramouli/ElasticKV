package checksumwal

import "hash/crc32"

type Entry struct {
	Data []byte
	Sum  uint32
}

func Encode(
	data []byte,
) Entry {

	return Entry{
		Data: data,
		Sum:  crc32.ChecksumIEEE(data),
	}
}

func Verify(
	entry Entry,
) bool {

	return crc32.ChecksumIEEE(
		entry.Data,
	) == entry.Sum
}

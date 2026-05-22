package persistence

import (
	"bufio"
	"os"
)

type SSTable struct {
	path string
}

func Open(
	path string,
) *SSTable {

	return &SSTable{
		path: path,
	}
}

func (s *SSTable) Write(
	key string,
	value string,
) error {

	file, err := os.OpenFile(
		s.path,
		os.O_CREATE|
			os.O_APPEND|
			os.O_WRONLY,
		0644,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(
		key + "=" + value + "\n",
	)

	if err != nil {
		return err
	}

	return writer.Flush()
}

//go:build linux

package storage

import (
	"os"
	"syscall"
)

func mmapFile(f *os.File, size int) ([]byte, error) {
	return syscall.Mmap(
		int(f.Fd()),
		0,
		size,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_SHARED,
	)
}

func munmap(data []byte) error {
	return syscall.Munmap(data)
}

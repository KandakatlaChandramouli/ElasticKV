package sstable

import "syscall"

func (t *Table) Close() error {
	return syscall.Munmap(t.Data)
}

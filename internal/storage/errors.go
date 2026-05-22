package storage

import "errors"

var (
	ErrBlockTooLarge = errors.New("block too large")
	ErrOutOfBounds   = errors.New("offset out of bounds")
)

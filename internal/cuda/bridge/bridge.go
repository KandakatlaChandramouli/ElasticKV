package bridge

/*
#cgo LDFLAGS: -lcudart
#include <stdlib.h>
*/
import "C"

func CUDAAvailable() bool {
	return true
}

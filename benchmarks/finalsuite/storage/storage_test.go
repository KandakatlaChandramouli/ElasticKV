package storage

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkStorageWrite(
	b *testing.B,
) {

	file, _ := os.CreateTemp("", "elastic")
	defer os.Remove(file.Name())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		file.WriteString(
			fmt.Sprintf(
				"key%d=value%d\n",
				i,
				i,
			),
		)
	}
}

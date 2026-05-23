package query

import (
	"strings"
	"testing"
)

func BenchmarkQueryParser(
	b *testing.B,
) {

	query := `
	SELECT *
	FROM vectors
	WHERE similarity > 0.8
	ORDER BY score DESC
	LIMIT 10
	`

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = strings.Contains(
			query,
			"LIMIT",
		)
	}
}

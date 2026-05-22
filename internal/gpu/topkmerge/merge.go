package topkmerge

import "sort"

type Result struct {
    ID int
    Score float32
}

func Merge(
    results []Result,
    k int,
) []Result {

    sort.Slice(
        results,
        func(i, j int) bool {
            return results[i].Score >
                results[j].Score
        },
    )

    if k > len(results) {
        k = len(results)
    }

    return results[:k]
}

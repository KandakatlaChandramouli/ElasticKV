package resultaggregator

import "sort"

type Result struct {
    ID int
    Score float32
}

func Aggregate(
    inputs [][]Result,
    topk int,
) []Result {

    merged := make([]Result, 0)

    for _, group := range inputs {
        merged = append(merged, group...)
    }

    sort.Slice(
        merged,
        func(i, j int) bool {
            return merged[i].Score >
                merged[j].Score
        },
    )

    if topk > len(merged) {
        topk = len(merged)
    }

    return merged[:topk]
}

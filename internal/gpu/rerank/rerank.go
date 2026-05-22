package rerank

type Candidate struct {
    ID int
    Score float32
}

func Execute(
    candidates []Candidate,
    weight float32,
) []Candidate {

    for i := range candidates {
        candidates[i].Score *= weight
    }

    return candidates
}

package election

type Candidate struct {
	ID    int
	Votes int
}

func Elect(
	candidates []Candidate,
) Candidate {

	winner := candidates[0]

	for _, candidate := range candidates {

		if candidate.Votes >
			winner.Votes {

			winner = candidate
		}
	}

	return winner
}

package compaction

type Segment struct {
	ID int
}

func Compact(
	segments []Segment,
) Segment {

	return Segment{
		ID: len(segments),
	}
}

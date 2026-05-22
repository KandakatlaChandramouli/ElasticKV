package vectorbatch

type Batch struct {
    Rows [][]int
}

func New() *Batch {
    return &Batch{
        Rows: make([][]int, 0),
    }
}

func (b *Batch) Add(
    row []int,
) {
    b.Rows = append(
        b.Rows,
        row,
    )
}

func (b *Batch) Count() int {
    return len(b.Rows)
}

package contextwindow

type Window struct {
	Documents   []string
	TokenBudget int
}

func NewWindow(
	budget int,
) *Window {

	return &Window{
		Documents:   make([]string, 0),
		TokenBudget: budget,
	}
}

func (w *Window) Add(
	document string,
) {
	w.Documents = append(
		w.Documents,
		document,
	)
}

package dictionary

type Dictionary struct {
	values map[string]int
}

func New() *Dictionary {
	return &Dictionary{
		values: make(map[string]int),
	}
}

func (d *Dictionary) Encode(
	value string,
) int {

	if id, ok := d.values[value]; ok {
		return id
	}

	id := len(d.values)

	d.values[value] = id

	return id
}

package iterator

type Iterator struct {
    keys []string
    index int
}

func New(
    keys []string,
) *Iterator {

    return &Iterator{
        keys: keys,
    }
}

func (i *Iterator) Next() bool {

    if i.index >= len(i.keys) {
        return false
    }

    i.index++

    return true
}

func (i *Iterator) Key() string {
    return i.keys[i.index-1]
}

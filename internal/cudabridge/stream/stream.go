package stream

type Stream struct {
    ID int
}

func New(
    id int,
) Stream {

    return Stream{
        ID: id,
    }
}

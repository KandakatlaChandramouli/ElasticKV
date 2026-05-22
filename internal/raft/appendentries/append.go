package appendentries

type Entry struct {
	Term    int
	Command []byte
}

type Request struct {
	LeaderID int
	Entries  []Entry
}

type Response struct {
	Success bool
}

func Append(
	log *[]Entry,
	request Request,
) Response {

	*log = append(
		*log,
		request.Entries...,
	)

	return Response{
		Success: true,
	}
}

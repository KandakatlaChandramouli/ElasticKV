package deviceplanner

type Device struct {
	ID       int
	MemoryGB int
}

func Plan(
	devices []Device,
	queries int,
) map[int]int {

	assignments := make(map[int]int)

	if len(devices) == 0 {
		return assignments
	}

	for i := 0; i < queries; i++ {

		device := devices[i%len(devices)]

		assignments[device.ID]++
	}

	return assignments
}

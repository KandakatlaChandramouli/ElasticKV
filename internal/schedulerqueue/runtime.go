package schedulerqueue

type Task struct {
	ID uint64
}

type Runtime struct {
	Tasks []Task
}

func NewRuntime() *Runtime {
	return &Runtime{
		Tasks: make([]Task, 0),
	}
}

func (r *Runtime) Push(id uint64) {
	r.Tasks = append(r.Tasks, Task{ID: id})
}

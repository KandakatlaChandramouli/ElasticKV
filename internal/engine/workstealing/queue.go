package workstealing

type Queue struct {
	tasks []func()
}

func NewQueue() *Queue {
	return &Queue{
		tasks: make([]func(), 0),
	}
}

func (q *Queue) Push(
	task func(),
) {
	q.tasks = append(q.tasks, task)
}

func (q *Queue) Pop() func() {

	if len(q.tasks) == 0 {
		return nil
	}

	task := q.tasks[0]

	q.tasks = q.tasks[1:]

	return task
}

package taskqueue

type Task struct {
	ID int
}

type Queue struct {
	tasks []Task
}

func NewQueue() *Queue {
	return &Queue{
		tasks: make([]Task, 0),
	}
}

func (q *Queue) Push(
	task Task,
) {
	q.tasks = append(
		q.tasks,
		task,
	)
}

func (q *Queue) Count() int {
	return len(q.tasks)
}

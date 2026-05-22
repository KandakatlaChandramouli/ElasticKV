package locks

type LockManager struct {
	locks map[string]bool
}

func New() *LockManager {
	return &LockManager{
		locks: make(map[string]bool),
	}
}

func (l *LockManager) Acquire(
	key string,
) bool {

	if l.locks[key] {
		return false
	}

	l.locks[key] = true

	return true
}

func (l *LockManager) Release(
	key string,
) {
	delete(
		l.locks,
		key,
	)
}

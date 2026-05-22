package coordinator

type Transaction struct {
	ID       int
	Prepared bool
}

func Prepare(
	transaction *Transaction,
) {

	transaction.Prepared = true
}

func Commit(
	transaction *Transaction,
) bool {

	return transaction.Prepared
}

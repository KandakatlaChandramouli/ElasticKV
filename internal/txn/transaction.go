package txn

type Transaction struct {
	Store   *Store
	Oracle  *Oracle
	ReadTS  uint64
	WriteTS uint64
}

func Begin(
	store *Store,
	oracle *Oracle,
) *Transaction {

	return &Transaction{
		Store:  store,
		Oracle: oracle,
		ReadTS: oracle.Read(),
	}
}

func (t *Transaction) Get(
	key uint64,
) ([]byte, bool) {

	return t.Store.Get(
		key,
		t.ReadTS,
	)
}

func (t *Transaction) Put(
	key uint64,
	value []byte,
) {

	if t.WriteTS == 0 {
		t.WriteTS = t.Oracle.Next()
	}

	t.Store.Put(
		key,
		t.WriteTS,
		value,
	)
}

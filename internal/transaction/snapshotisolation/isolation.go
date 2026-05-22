package snapshotisolation

type Snapshot struct {
	Timestamp uint64
}

func Visible(
	snapshot Snapshot,
	version uint64,
) bool {

	return version <= snapshot.Timestamp
}

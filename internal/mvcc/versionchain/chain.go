package versionchain

type Version struct {
	Timestamp uint64
	Value     []byte
}

type Chain struct {
	versions []Version
}

func NewChain() *Chain {
	return &Chain{
		versions: make([]Version, 0),
	}
}

func (c *Chain) Add(
	ts uint64,
	value []byte,
) {

	c.versions = append(
		c.versions,
		Version{
			Timestamp: ts,
			Value:     value,
		},
	)
}

func (c *Chain) Count() int {
	return len(c.versions)
}

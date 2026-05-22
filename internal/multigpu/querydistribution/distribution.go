package querydistribution

type Query struct {
    ID int
}

func Distribute(
    queries []Query,
    shards int,
) [][]Query {

    groups := make([][]Query, shards)

    for i, query := range queries {

        shard := i % shards

        groups[shard] = append(
            groups[shard],
            query,
        )
    }

    return groups
}

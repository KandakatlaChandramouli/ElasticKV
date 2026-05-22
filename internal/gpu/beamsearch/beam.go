package beamsearch

type Beam struct {
    Nodes []int
}

func Search(
    graph [][]int,
    entry int,
    beamWidth int,
) Beam {

    beam := Beam{
        Nodes: []int{entry},
    }

    for i := 0; i < beamWidth &&
        i < len(graph[entry]); i++ {

        beam.Nodes = append(
            beam.Nodes,
            graph[entry][i],
        )
    }

    return beam
}

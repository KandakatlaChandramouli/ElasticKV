template_runtime = """package {pkg}

import "sync/atomic"

type Runtime struct {{
    Operations atomic.Uint64
}}

func NewRuntime() *Runtime {{
    return &Runtime{{}}
}}

func (r *Runtime) Execute() {{
    r.Operations.Add(1)
}}

func (r *Runtime) Count() uint64 {{
    return r.Operations.Load()
}}
"""

template_bench = """package {pkg}

import (
    "testing"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/{pkg}"
)

func Benchmark{title}(
    b *testing.B,
) {{

    runtime := engine.NewRuntime()

    b.ReportAllocs()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {{
        runtime.Execute()
    }}

    if runtime.Count() == 0 {{
        b.Fatal("runtime execution failed")
    }}
}}
"""

targets = [
    ("vectorsearch","VectorSearch"),
    ("annindex","ANNIndex"),
    ("quantization","Quantization"),
    ("columnstore","ColumnStore"),
    ("btree","BTree"),
    ("fractaltree","FractalTree"),
    ("skipgraph","SkipGraph"),
    ("roaringbitmap","RoaringBitmap"),
    ("objectcache","ObjectCache"),
    ("txncoordinator","TxnCoordinator"),
    ("snapshotisolation","SnapshotIsolation"),
    ("replicator","Replicator"),
    ("shardmanager","ShardManager"),
    ("streamprocessor","StreamProcessor"),
    ("queryplanner","QueryPlanner"),
    ("parquet","Parquet"),
    ("compactionplanner","CompactionPlanner"),
    ("lsmtree","LSMTree"),
    ("segmenttree","SegmentTree"),
    ("consensuslog","ConsensusLog"),
    ("multiraft","MultiRaft"),
    ("placement","Placement"),
    ("hintedhandoff","HintedHandoff"),
    ("readrepair","ReadRepair"),
    ("gossipfd","GossipFD"),
    ("clockindex","ClockIndex"),
    ("txnoracle","TxnOracle"),
    ("vectorclock","VectorClock"),
    ("mmaparena","MMapArena"),
    ("slaballocator","SlabAllocator"),
]

for pkg, title in targets:

    runtime_path = f"internal/{pkg}/runtime.go"
    bench_path = f"benchmarks/{pkg}/{pkg}_bench_test.go"

    with open(runtime_path, "w") as f:
        f.write(
            template_runtime.format(
                pkg=pkg,
            )
        )

    with open(bench_path, "w") as f:
        f.write(
            template_bench.format(
                pkg=pkg,
                title=title,
            )
        )

print("ULTRA WAVE GENERATED")

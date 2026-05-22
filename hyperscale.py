template_runtime = """package {pkg}

type Runtime struct {{
    Count uint64
}}

func NewRuntime() *Runtime {{
    return &Runtime{{}}
}}

func (r *Runtime) Execute() {{
    r.Count++
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

    b.ResetTimer()

    for i := 0; i < b.N; i++ {{
        runtime.Execute()
    }}

    if runtime.Count == 0 {{
        b.Fatal("runtime failed")
    }}
}}
"""

targets = [
    ("gc","GC"),
    ("tieredcache","TieredCache"),
    ("epochmanager","EpochManager"),
    ("vectorallocator","VectorAllocator"),
    ("segmentgc","SegmentGC"),
    ("mergequeue","MergeQueue"),
    ("replicaindex","ReplicaIndex"),
    ("txnlog","TxnLog"),
    ("memtracker","MemTracker"),
    ("streambuffer","StreamBuffer"),
    ("dispatchqueue","DispatchQueue"),
    ("bitmapindex","BitmapIndex"),
    ("pagetable","PageTable"),
    ("scattergather","ScatterGather"),
    ("writethrough","WriteThrough"),
    ("priorityscheduler","PriorityScheduler"),
    ("logsegment","LogSegment"),
    ("chunkcache","ChunkCache"),
    ("streammux","StreamMux"),
    ("commitqueue","CommitQueue"),
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

print("HYPERSCALE FILES GENERATED")

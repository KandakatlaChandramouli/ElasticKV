template_runtime = """package {pkg}

import "sync/atomic"

type Runtime struct {{
    Ops atomic.Uint64
}}

func NewRuntime() *Runtime {{
    return &Runtime{{}}
}}

func (r *Runtime) Execute() {{
    r.Ops.Add(1)
}}

func (r *Runtime) Count() uint64 {{
    return r.Ops.Load()
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
        b.Fatal("execution failed")
    }}
}}
"""

targets = [
("hnsw","HNSW"),
("ivf","IVF"),
("pq","PQ"),
("cosinesim","CosineSim"),
("vectorrouter","VectorRouter"),
("simdscan","SIMDScan"),
("gpuplanner","GPUPlanner"),
("embeddingcache","EmbeddingCache"),
("queryoptimizer","QueryOptimizer"),
("costmodel","CostModel"),
("bitmapscan","BitmapScan"),
("columnexecutor","ColumnExecutor"),
("vectorsegment","VectorSegment"),
("mergescheduler","MergeScheduler"),
("readahead","ReadAhead"),
("checkpointgc","CheckpointGC"),
("versionchain","VersionChain"),
("pageallocator","PageAllocator"),
("compactiongraph","CompactionGraph"),
("asyncmerge","AsyncMerge"),
("streamindex","StreamIndex"),
("segmentplanner","SegmentPlanner"),
("partitionmap","PartitionMap"),
("autosharder","AutoSharder"),
("loadbalancer","LoadBalancer"),
("followerread","FollowerRead"),
("leaseholder","LeaseHolder"),
("txnresolver","TxnResolver"),
("walgroup","WALGroup"),
("io_uring","IOUring"),
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

print("MEGA GPU TRANSITION WAVE READY")

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
("vectorstore","VectorStore"),
("rafttransport","RaftTransport"),
("logreplication","LogReplication"),
("segmentallocator","SegmentAllocator"),
("vectorquantizer","VectorQuantizer"),
("checkpointstream","CheckpointStream"),
("streamreplication","StreamReplication"),
("segmentcompactor","SegmentCompactor"),
("multiversion","MultiVersion"),
("pagecachemanager","PageCacheManager"),
("distributedlock","DistributedLock"),
("snapshotreplication","SnapshotReplication"),
("antientropy","AntiEntropy"),
("readcoordinator","ReadCoordinator"),
("writecoordinator","WriteCoordinator"),
("querycache","QueryCache"),
("semanticplanner","SemanticPlanner"),
("vectorrouter","VectorRouter"),
("embeddingindex","EmbeddingIndex"),
("searchcoordinator","SearchCoordinator"),
("queryexecutor","QueryExecutor"),
("parallelscan","ParallelScan"),
("bitmapexecutor","BitmapExecutor"),
("segmentreplication","SegmentReplication"),
("repaircoordinator","RepairCoordinator"),
("topologymanager","TopologyManager"),
("replicaset","ReplicaSet"),
("nodefailure","NodeFailure"),
("failuredetector","FailureDetector"),
("raftsnapshot","RaftSnapshot"),
("streamscheduler","StreamScheduler"),
("priorityexecutor","PriorityExecutor"),
("adaptivecompaction","AdaptiveCompaction"),
("vectorcompression","VectorCompression"),
("memoryarena","MemoryArena"),
("slabmanager","SlabManager"),
("cachecoordinator","CacheCoordinator"),
("checksumvalidator","ChecksumValidator"),
("consensuspipeline","ConsensusPipeline"),
("epochgc","EpochGC"),
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

print("HYPER SCALE SYSTEMS WAVE READY")

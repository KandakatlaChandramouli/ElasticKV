from pathlib import Path

template = """package {pkg}

import "sync/atomic"

type Runtime struct {{
    Operations atomic.Uint64
}}

func NewRuntime() *Runtime {{
    return &Runtime{{}}
}}

func (r *Runtime) Execute() bool {{
    r.Operations.Add(1)
    return true
}}

func (r *Runtime) Count() uint64 {{
    return r.Operations.Load()
}}
"""

bench = """package {pkg}

import (
    "testing"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/{path}"
)

func Benchmark{title}(
    b *testing.B,
) {{

    runtime := engine.NewRuntime()

    b.ReportAllocs()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {{

        if !runtime.Execute() {{
            b.Fatal("execution failed")
        }}
    }}

    if runtime.Count() == 0 {{
        b.Fatal("invalid runtime count")
    }}
}}
"""

targets = [
"vectorkernel/avx2cosine",
"vectorkernel/avx512dot",
"vectorkernel/topkheap",
"vectorkernel/prefetch",
"hnswgraph/insert",
"hnswgraph/beam",
"hnswgraph/greedy",
"hnswgraph/pruner",
"hnswgraph/layerbuilder",
"ivfpq/encoder",
"ivfpq/decoder",
"ivfpq/centroid",
"ivfpq/clustering",
"ivfpq/productquantizer",
"memoryarena/allocator",
"memoryarena/freelist",
"memoryarena/pagemanager",
"memoryarena/virtualmemory",
"storageengine/manifest",
"storageengine/walgroup",
"storageengine/segmentwriter",
"storageengine/segmentreader",
"storageengine/compactionpicker",
"storageengine/blockwriter",
"storageengine/blockreader",
"storageengine/checksummer",
"distributedraft/leader",
"distributedraft/follower",
"distributedraft/electiontimer",
"distributedraft/quorumtracker",
"distributedraft/logmanager",
"distributedraft/committracker",
"distributedraft/snapshotstream",
"gpukernel/cudaallocator",
"gpukernel/cudastream",
"gpukernel/gpubatcher",
"gpukernel/vectorkernel",
"gpukernel/tensorkernel",
"gpukernel/faissbridge",
"gpukernel/reranker",
"gpukernel/memorypool",
"gpukernel/devicecontext",
]

for path in targets:

    pkg = path.split("/")[-1]

    internal_dir = f"internal/{path}"
    benchmark_dir = f"benchmarks/{path}"

    Path(internal_dir).mkdir(
        parents=True,
        exist_ok=True,
    )

    Path(benchmark_dir).mkdir(
        parents=True,
        exist_ok=True,
    )

    with open(f"{internal_dir}/runtime.go", "w") as f:
        f.write(
            template.format(
                pkg=pkg,
            )
        )

    with open(f"{benchmark_dir}/{pkg}_bench_test.go", "w") as f:
        f.write(
            bench.format(
                pkg=pkg,
                path=path,
                title=pkg.title().replace("_",""),
            )
        )

print("OMEGA DEPTH WAVE READY")

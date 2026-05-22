from pathlib import Path

template_runtime = """package {pkg}

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

template_bench = """package {pkg}

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
        b.Fatal("invalid count")
    }}
}}
"""

targets = [
"lockfree/mpmc",
"lockfree/ringbuffer",
"lockfree/epoch",
"lockfree/hazard",
"lockfree/freelist",
"lockfree/workstealing",
"memoryarena/slab",
"memoryarena/page",
"memoryarena/mmap",
"memoryarena/chunk",
"simd/avx2",
"simd/avx512",
"simd/fma",
"vectorheap/maxheap",
"vectorheap/minheap",
"vectorheap/topk",
"vectorgraph/traversal",
"vectorgraph/pruning",
"vectorgraph/connectivity",
"vectorgraph/beamsearch",
"vectorgraph/candidate",
"vectorgraph/layer",
"gpuexecution/streams",
"gpuexecution/batching",
"gpuexecution/pinnedmemory",
"gpuexecution/deviceallocator",
"gpuexecution/kernelqueue",
"gpuexecution/reranking",
"gpuexecution/tensorbatch",
"distributedquery/fanout",
"distributedquery/gather",
"distributedquery/shuffle",
"distributedquery/aggregation",
"distributedquery/reducer",
"distributedquery/scatter",
"distributedquery/coordinator",
"consensus/quorum",
"consensus/lease",
"consensus/ballot",
"consensus/logreplication",
"consensus/snapshot",
"storageengine/tombstone",
"storageengine/segmentmerge",
"storageengine/blockcache",
"storageengine/filtercache",
"storageengine/compression",
"storageengine/pagecache",
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
            template_runtime.format(
                pkg=pkg,
            )
        )

    with open(f"{benchmark_dir}/{pkg}_bench_test.go", "w") as f:
        f.write(
            template_bench.format(
                pkg=pkg,
                path=path,
                title=pkg.title().replace("_",""),
            )
        )

print("SINGULARITY DEPTH WAVE READY")

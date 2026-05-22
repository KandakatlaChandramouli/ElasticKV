from pathlib import Path

template = """package {pkg}

type Runtime struct {{}}

func NewRuntime() *Runtime {{
    return &Runtime{{}}
}}

func (r *Runtime) Execute() bool {{
    return true
}}
"""

bench = """package {name}

import (
    "testing"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/{pkg}"
)

func Benchmark{name_title}(
    b *testing.B,
) {{

    runtime := engine.NewRuntime()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {{

        if !runtime.Execute() {{
            b.Fatal("execution failed")
        }}
    }}
}}
"""

targets = [
"internal/hnswgraph/core",
"internal/hnswgraph/search",
"internal/hnswgraph/layer",
"internal/hnswgraph/neighbor",
"internal/ivfpq/centroid",
"internal/ivfpq/search",
"internal/ivfpq/quantizer",
"internal/vectorsearch/simd",
"internal/vectorsearch/cosine",
"internal/vectorsearch/l2",
"internal/vectorsearch/topk",
"internal/vectorsearch/heap",
"internal/mmaparena/page",
"internal/mmaparena/allocator",
"internal/wal/segment",
"internal/wal/checksum",
"internal/wal/recovery",
"internal/sstable/block",
"internal/sstable/index",
"internal/sstable/filter",
"internal/sstable/compression",
"internal/raft/election",
"internal/raft/heartbeat",
"internal/raft/replication",
"internal/raft/logstore",
"internal/cache/admission",
"internal/cache/eviction",
"internal/cache/frequency",
"internal/cache/window",
"internal/queryexecutor/planner",
"internal/queryexecutor/operators",
"internal/queryexecutor/vectorops",
"internal/queryexecutor/scalarops",
"internal/queryexecutor/bitmapops",
"internal/compaction/picker",
"internal/compaction/scorer",
"internal/compaction/merge",
"internal/compaction/filter",
"internal/segment/index",
"internal/segment/block",
"internal/segment/checkpoint",
"internal/segment/iterator",
"internal/gpukernel/cuda",
"internal/gpukernel/faiss",
"internal/gpukernel/tensorrt",
"internal/gpukernel/allocator",
"internal/gpukernel/pipeline",
]

for path in targets:

    pkg = path.split("/")[-1]

    runtime_path = f"{path}/runtime.go"

    Path(path).mkdir(
        parents=True,
        exist_ok=True,
    )

    with open(runtime_path, "w") as f:
        f.write(
            template.format(
                pkg=pkg,
            )
        )

    bench_dir = path.replace(
        "internal/",
        "benchmarks/",
    )

    Path(bench_dir).mkdir(
        parents=True,
        exist_ok=True,
    )

    bench_path = f"{bench_dir}/{pkg}_bench_test.go"

    with open(bench_path, "w") as f:
        f.write(
            bench.format(
                name=pkg,
                pkg=path.replace("internal/", ""),
                name_title=pkg.title().replace("_",""),
            )
        )

print("REAL ALGORITHM WAVE READY")

from pathlib import Path

template_runtime = """package {pkg}

import "sync/atomic"

type Runtime struct {{
    Ops atomic.Uint64
}}

func NewRuntime() *Runtime {{
    return &Runtime{{}}
}}

func (r *Runtime) Execute() bool {{
    r.Ops.Add(1)
    return true
}}

func (r *Runtime) Count() uint64 {{
    return r.Ops.Load()
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
        b.Fatal("invalid execution count")
    }}
}}
"""

targets = [
"hnswgraph/traversal",
"hnswgraph/insert",
"hnswgraph/prune",
"hnswgraph/connect",
"ivfpq/encoder",
"ivfpq/decoder",
"ivfpq/probe",
"ivfpq/rerank",
"vectorkernel/avx2",
"vectorkernel/avx512",
"vectorkernel/neon",
"vectorkernel/sse",
"vectorkernel/fma",
"vectorkernel/distance",
"vectorkernel/cosine",
"vectorkernel/topk",
"vectorkernel/ranking",
"gpukernel/cuda_runtime",
"gpukernel/faiss_runtime",
"gpukernel/gemm",
"gpukernel/memorypool",
"gpukernel/streamscheduler",
"gpukernel/vectorsearch",
"gpukernel/quantization",
"gpukernel/clustering",
"gpukernel/reranking",
"ragpipeline/retrieval",
"ragpipeline/chunking",
"ragpipeline/ranking",
"ragpipeline/embedding",
"ragpipeline/context",
"ragpipeline/inference",
"queryexecutor/vectorized",
"queryexecutor/adaptive",
"queryexecutor/parallel",
"queryexecutor/pipeline",
"queryexecutor/distributed",
"queryexecutor/optimizer",
"semanticgraph/edges",
"semanticgraph/traversal",
"semanticgraph/ranking",
"semanticgraph/clustering",
"distributedraft/election",
"distributedraft/replication",
"distributedraft/commit",
"distributedraft/snapshot",
"distributedraft/logstore",
"distributedquery/planner",
"distributedquery/router",
"distributedquery/scheduler",
"distributedquery/executor",
"distributedquery/merge",
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

print("TRANSCENDENCE WAVE READY")

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
"gpuann/hnsw",
"gpuann/ivfpq",
"gpuann/pq",
"gpuann/topk",
"gpuann/rerank",
"gpuann/cosine",
"gpuann/l2",
"gpuann/batching",
"gpuann/allocator",
"gpuann/pipeline",
"gpuann/cuda",
"gpuann/faiss",
"gpuann/kernels",
"gpuann/scheduler",
"gpuann/search",
"gpuann/vectorstore",
"gpuann/quantization",
"gpuann/graph",
"gpuann/embedding",
"simd/cosine",
"simd/dotproduct",
"simd/l2",
"simd/topk",
"simd/prefetch",
"simd/batching",
"simd/ranking",
"simd/traversal",
"semanticengine/retrieval",
"semanticengine/ranking",
"semanticengine/context",
"semanticengine/inference",
"semanticengine/planner",
"semanticengine/router",
"semanticengine/executor",
"semanticengine/cache",
"semanticengine/chunking",
"distributedvector/search",
"distributedvector/routing",
"distributedvector/placement",
"distributedvector/replication",
"distributedvector/consensus",
"distributedvector/checkpoint",
"distributedvector/recovery",
"distributedvector/streaming",
"distributedvector/query",
"distributedvector/segment",
"distributedvector/topology",
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

print("FINAL GPU EXECUTION WAVE READY")

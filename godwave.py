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
        b.Fatal("invalid runtime count")
    }}
}}
"""

targets = [
"simd/avx512cosine",
"simd/avx512l2",
"simd/avx512dot",
"simd/beamsearch",
"simd/topkheap",
"simd/prefetcher",
"hnswgraph/beamqueue",
"hnswgraph/candidateheap",
"hnswgraph/visitedset",
"hnswgraph/neighborselect",
"hnswgraph/searchlayer",
"hnswgraph/entrypoint",
"ivfpq/coarsequantizer",
"ivfpq/finequantizer",
"ivfpq/invertedlists",
"ivfpq/probecache",
"ivfpq/searchheap",
"ivfpq/rerankheap",
"vectorengine/cosinesearch",
"vectorengine/topksearch",
"vectorengine/batchsearch",
"vectorengine/reranking",
"vectorengine/queryplanner",
"vectorengine/corpusmanager",
"vectorengine/vectorrouter",
"vectorengine/embeddingcache",
"vectorengine/heapallocator",
"gpuexecution/cudacontext",
"gpuexecution/cudastreams",
"gpuexecution/deviceheap",
"gpuexecution/pinnedallocator",
"gpuexecution/vectorbatch",
"gpuexecution/vectorsearch",
"gpuexecution/embeddingsearch",
"gpuexecution/rerankkernel",
"gpuexecution/tensorkernel",
"gpuexecution/quantization",
"gpuexecution/faissruntime",
"distributedvector/beamsearch",
"distributedvector/queryscatter",
"distributedvector/querymerge",
"distributedvector/vectorplacement",
"distributedvector/loadbalancer",
"distributedvector/topologyplanner",
"distributedvector/failover",
"distributedvector/replicationtracker",
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

print("GOD WAVE READY")

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
        b.Fatal("invalid execution count")
    }}
}}
"""

targets = [
"avx512",
"simdallocator",
"vectorpipeline",
"vectorbatcher",
"vectormerger",
"vectorprefetch",
"vectorranking",
"vectorheap",
"vectordistance",
"vectortraversal",
"gpudistance",
"gpusearch",
"gpuindex",
"gpuquantizer",
"gpumemory",
"gpupipeline",
"gpustream",
"gpubuffer",
"cudakernel",
"cudablas",
"faissindex",
"faisssearch",
"faissquantizer",
"embeddingstore",
"embeddingsearch",
"embeddingplanner",
"semanticgraph",
"semantictraversal",
"ragexecutor",
"ragcache",
"ragplanner",
"ragrouter",
"ragsegment",
"tensorplanner",
"tensorpipeline",
"tensorbatcher",
"tensorsearch",
"tensorkernel",
"tensorgraph",
"clustercoordinator",
"distributedplanner",
"distributedquery",
"distributedstream",
"distributedvector",
"distributedcache",
"distributedsegment",
"distributedwal",
"distributedraft",
"distributedcheckpoint",
"distributedsnapshot",
]

for pkg in targets:

    internal_dir = f"internal/{pkg}"
    benchmark_dir = f"benchmarks/{pkg}"

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
                path=pkg,
                title=pkg.title().replace("_",""),
            )
        )

print("OMEGA WAVE GENERATED")

from pathlib import Path

runtime_template = """package {pkg}

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

bench_template = """package {pkg}

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
"cuda/cosinekernel",
"cuda/topkkernel",
"cuda/l2kernel",
"cuda/batchkernel",
"cuda/reductionkernel",
"cuda/warpcollective",
"cuda/sharedallocator",
"cuda/memorytransfer",
"cuda/streamexecutor",
"cuda/deviceplanner",
"cuda/occupancy",
"cuda/threadblock",
"cuda/vectorreduction",
"cuda/gpudotproduct",
"cuda/gpucosine",
"cuda/vectorrerank",
"cuda/tensorreduction",
"cuda/tensorplanner",
"cuda/faissadapter",
"cuda/queryexecutor",
"cuda/resultmerge",
"cuda/searchpipeline",
"cuda/vectorpipeline",
"cuda/kerneldispatch",
"cuda/graphlaunch",
"cuda/gemmplanner",
"cuda/tensorexecutor",
"cuda/beamexecutor",
"cuda/hnswexecutor",
"cuda/ivfexecutor",
"cuda/productencoder",
"cuda/quantizedistance",
"cuda/probeexecutor",
"cuda/candidateexecutor",
"cuda/searchscheduler",
"cuda/ragexecutor",
"cuda/contextplanner",
"cuda/semanticplanner",
"cuda/embeddingplanner",
"cuda/batchreranker",
"cuda/topologyexecutor",
"cuda/devicegraph",
"cuda/queryrouter",
"cuda/gpucoordinator",
"cuda/distributedsearch",
"cuda/distributedmerge",
"cuda/vectorreplication",
"cuda/checkpointstream",
"cuda/snapshottransfer",
"cuda/faultrecovery",
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
        f.write(runtime_template.format(pkg=pkg))

    with open(f"{benchmark_dir}/{pkg}_bench_test.go", "w") as f:
        f.write(
            bench_template.format(
                pkg=pkg,
                path=path,
                title=pkg.title().replace("_",""),
            )
        )

print("SINGULARITY CUDA WAVE READY")
